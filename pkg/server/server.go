package server

import (
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.ibm.com/blockchaindb/library/pkg/constants"
	"github.ibm.com/blockchaindb/library/pkg/logger"
	"github.ibm.com/blockchaindb/protos/types"
	"github.ibm.com/blockchaindb/server/config"
	"github.ibm.com/blockchaindb/server/pkg/blockstore"
	"github.ibm.com/blockchaindb/server/pkg/fileops"
	"github.ibm.com/blockchaindb/server/pkg/worldstate/leveldb"
)

// DBAndHTTPServer holds the database and http server objects
type DBAndHTTPServer struct {
	dbServ  *dbServer
	handler http.Handler
	listen  net.Listener
	conf    *config.Configurations
	logger  *logger.SugarLogger
}

// New creates a object of DBAndHTTPServer
func New(conf *config.Configurations) (*DBAndHTTPServer, error) {
	c := &logger.Config{
		Level:         conf.Node.LogLevel,
		OutputPath:    []string{"stdout"},
		ErrOutputPath: []string{"stderr"},
		Encoding:      "console",
	}
	logger, err := logger.New(c)
	if err != nil {
		return nil, err
	}

	dbServ, err := newDBServer(conf, logger)
	if err != nil {
		return nil, errors.Wrap(err, "error while creating the database object")
	}

	router := mux.NewRouter()
	router.HandleFunc(constants.GetDBStatus, dbServ.handleDBStatusQuery).Methods(http.MethodGet)
	router.HandleFunc(constants.GetData, dbServ.handleDataQuery).Methods(http.MethodGet)
	router.HandleFunc(constants.GetUser, dbServ.handleUserQuery).Methods(http.MethodGet)
	router.HandleFunc(constants.GetConfig, dbServ.handleConfigQuery).Methods(http.MethodGet)
	router.HandleFunc(constants.PostDataTx, dbServ.handleDataTransaction).Methods(http.MethodPost)
	router.HandleFunc(constants.PostUserTx, dbServ.handleUserAdminTransaction).Methods(http.MethodPost)
	router.HandleFunc(constants.PostDBTx, dbServ.handleDBAdminTransaction).Methods(http.MethodPost)
	router.HandleFunc(constants.PostConfigTx, dbServ.handleConfigTransaction).Methods(http.MethodPost)

	netConf := conf.Node.Network
	addr := fmt.Sprintf("%s:%d", netConf.Address, netConf.Port)
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, errors.Wrap(err, "error while creating a tcp listener")
	}

	return &DBAndHTTPServer{
		dbServ:  dbServ,
		handler: router,
		listen:  listen,
		conf:    conf,
		logger:  logger,
	}, nil
}

// Start starts the server
func (s *DBAndHTTPServer) Start() error {
	blockHeight, err := s.dbServ.blockStore.Height()
	if err != nil {
		return err
	}
	if blockHeight == 0 {
		if err := s.dbServ.prepareAndCommitConfigTx(s.conf); err != nil {
			return errors.Wrap(err, "error while preparing and committing config transaction")
		}
	}

	s.logger.Infof("Starting the server on %s", s.listen.Addr().String())

	go func() {
		if err := http.Serve(s.listen, s.handler); err != nil {
			switch err.(type) {
			case *net.OpError:
				s.logger.Warn("network connection is closed")
			default:
				s.logger.Panicf("server stopped unexpectedly, %v", err)
			}
		}
	}()

	return nil
}

// Stop stops the server
func (s *DBAndHTTPServer) Stop() error {
	if s == nil || s.listen == nil {
		return nil
	}

	s.logger.Infof("Stopping the server listening on %s\n", s.listen.Addr().String())
	if err := s.listen.Close(); err != nil {
		return errors.Wrap(err, "error while closing the network listener")
	}

	return s.dbServ.close()
}

type dbServer struct {
	*queryProcessor
	*transactionProcessor
	logger *logger.SugarLogger
}

func newDBServer(conf *config.Configurations, logger *logger.SugarLogger) (*dbServer, error) {
	if conf.Node.Database.Name != "leveldb" {
		return nil, errors.New("only leveldb is supported as the state database")
	}

	ledgerDir := conf.Node.Database.LedgerDirectory
	if err := createLedgerDir(ledgerDir); err != nil {
		return nil, err
	}

	levelDB, err := leveldb.Open(
		&leveldb.Config{
			DBRootDir: constructWorldStatePath(ledgerDir),
			Logger:    logger,
		},
	)
	if err != nil {
		return nil, errors.WithMessage(err, "error while creating the world state database")
	}

	blockStore, err := blockstore.Open(
		&blockstore.Config{
			StoreDir: constructBlockStorePath(ledgerDir),
			Logger:   logger,
		},
	)
	if err != nil {
		return nil, errors.WithMessage(err, "error while creating the block store")
	}

	blkHeight, err := blockStore.Height()
	if err != nil {
		return nil, errors.WithMessage(err, "error while fetching the block store height")
	}

	qProcConfig := &queryProcessorConfig{
		nodeID:     conf.Node.Identity.ID,
		db:         levelDB,
		blockStore: blockStore,
		logger:     logger,
	}

	txProcConf := &txProcessorConfig{
		db:                 levelDB,
		blockStore:         blockStore,
		blockHeight:        blkHeight,
		txQueueLength:      conf.Node.QueueLength.Transaction,
		txBatchQueueLength: conf.Node.QueueLength.ReorderedTransactionBatch,
		blockQueueLength:   conf.Node.QueueLength.Block,
		maxTxCountPerBatch: conf.Consensus.MaxTransactionCountPerBlock,
		batchTimeout:       conf.Consensus.BlockTimeout,
		logger:             logger,
	}

	return &dbServer{
		queryProcessor:       newQueryProcessor(qProcConfig),
		transactionProcessor: newTransactionProcessor(txProcConf),
		logger:               logger,
	}, nil
}

func (db *dbServer) close() error {
	if err := db.queryProcessor.close(); err != nil {
		return err
	}

	return db.transactionProcessor.close()
}

func (db *dbServer) handleDBStatusQuery(w http.ResponseWriter, r *http.Request) {
	_, _, composedErr := db.preProcessQuery(w, r)
	if composedErr {
		return
	}

	params := mux.Vars(r)
	dbName, ok := params["dbname"]
	if !ok {
		SendHTTPResponse(w, http.StatusBadRequest, "query error - bad or missing database name")
		return
	}

	//TODO: verify signature

	dbStatus, err := db.getDBStatus(dbName)
	if err != nil {
		SendHTTPResponse(
			w,
			http.StatusInternalServerError,
			&ResponseErr{
				Error: "error while processing [" + r.URL.String() + "] because " + err.Error(),
			},
		)
		return
	}

	SendHTTPResponse(w, http.StatusOK, dbStatus)
}

func (db *dbServer) handleDataQuery(w http.ResponseWriter, r *http.Request) {
	userID, _, composedErr := db.preProcessQuery(w, r)
	if composedErr {
		return
	}

	params := mux.Vars(r)
	dbName, ok := params["dbname"]
	if !ok {
		SendHTTPResponse(w,
			http.StatusBadRequest,
			&ResponseErr{
				Error: "query error - bad or missing database name",
			})
		return
	}
	key, ok := params["key"]
	if !ok {
		SendHTTPResponse(w,
			http.StatusBadRequest,
			&ResponseErr{
				Error: "query error - bad or missing key",
			})
		return
	}

	//TODO: verify signature

	data, err := db.getData(dbName, userID, key)
	if err != nil {
		var status int

		switch err.(type) {
		case *permissionErr:
			status = http.StatusForbidden
		default:
			status = http.StatusInternalServerError
		}

		SendHTTPResponse(
			w,
			status,
			&ResponseErr{
				Error: "error while processing [" + r.URL.String() + "] because " + err.Error(),
			})
		return
	}

	SendHTTPResponse(w, http.StatusOK, data)
}

func (db *dbServer) handleUserQuery(w http.ResponseWriter, r *http.Request) {
	querierUserID, _, composedErr := db.preProcessQuery(w, r)
	if composedErr {
		return
	}

	params := mux.Vars(r)
	targetUserID, ok := params["userid"]
	if !ok {
		SendHTTPResponse(w, http.StatusBadRequest, &ResponseErr{"query error - bad or missing userid"})
		return
	}

	//TODO: verify signature

	user, err := db.getUser(querierUserID, targetUserID)
	if err != nil {
		var status int

		switch err.(type) {
		case *permissionErr:
			status = http.StatusForbidden
		default:
			status = http.StatusInternalServerError
		}

		SendHTTPResponse(
			w,
			status,
			&ResponseErr{"error while processing [" + r.URL.String() + "] because " + err.Error()},
		)
		return
	}

	SendHTTPResponse(w, http.StatusOK, user)
}

func (db *dbServer) handleConfigQuery(w http.ResponseWriter, r *http.Request) {
	_, _, composedErr := db.preProcessQuery(w, r)
	if composedErr {
		return
	}

	config, err := db.getConfig()
	if err != nil {
		SendHTTPResponse(
			w,
			http.StatusInternalServerError,
			&ResponseErr{"error while processing [" + r.URL.String() + "] because " + err.Error()},
		)
		return
	}

	SendHTTPResponse(w, http.StatusOK, config)
}

func (db *dbServer) preProcessQuery(w http.ResponseWriter, r *http.Request) (string, []byte, bool) {
	composedErr := true

	querierUserID, signature, err := validateAndParseHeader(&r.Header)
	if err != nil {
		SendHTTPResponse(w, http.StatusBadRequest, &ResponseErr{err.Error()})
		return "", nil, composedErr
	}

	exist, err := db.identityQuerier.DoesUserExist(querierUserID)
	if err != nil {
		SendHTTPResponse(w, http.StatusInternalServerError, &ResponseErr{err.Error()})
		return "", nil, composedErr
	}
	if !exist {
		SendHTTPResponse(
			w,
			http.StatusForbidden,
			&ResponseErr{
				Error: r.URL.String() + " query is rejected as the submitting user [" + querierUserID + "] does not exist in the cluster",
			},
		)
		return "", nil, composedErr
	}

	return querierUserID, signature, !composedErr
}

func (db *dbServer) handleDataTransaction(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	tx := &types.DataTxEnvelope{}
	if err := d.Decode(tx); err != nil {
		SendHTTPResponse(w, http.StatusBadRequest, &ResponseErr{err.Error()})
		return
	}

	// TODO: verify signature
	db.handleTransaction(w, tx.Payload.UserID, tx)
}

func (db *dbServer) handleUserAdminTransaction(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	tx := &types.UserAdministrationTxEnvelope{}
	if err := d.Decode(tx); err != nil {
		SendHTTPResponse(w, http.StatusBadRequest, &ResponseErr{err.Error()})
		return
	}

	// TODO: verify signature
	db.handleTransaction(w, tx.Payload.UserID, tx)
}

func (db *dbServer) handleDBAdminTransaction(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	tx := &types.DBAdministrationTxEnvelope{}
	if err := d.Decode(tx); err != nil {
		SendHTTPResponse(w, http.StatusBadRequest, &ResponseErr{err.Error()})
		return
	}

	// TODO: verify signature
	db.handleTransaction(w, tx.Payload.UserID, tx)
}

func (db *dbServer) handleConfigTransaction(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	tx := &types.ConfigTxEnvelope{}
	if err := d.Decode(tx); err != nil {
		SendHTTPResponse(w, http.StatusBadRequest, &ResponseErr{err.Error()})
		return
	}

	// TODO: verify signature
	db.handleTransaction(w, tx.Payload.UserID, tx)
}

func (db *dbServer) handleTransaction(w http.ResponseWriter, userID string, tx interface{}) {
	exist, err := db.identityQuerier.DoesUserExist(userID)
	if err != nil {
		SendHTTPResponse(w, http.StatusBadRequest, &ResponseErr{err.Error()})
		return
	}
	if !exist {
		SendHTTPResponse(
			w,
			http.StatusForbidden,
			&ResponseErr{
				Error: "transaction is rejected as the submitting user [" + userID + "] does not exist in the cluster",
			},
		)
		return
	}

	if err := db.submitTransaction(tx); err != nil {
		SendHTTPResponse(w, http.StatusInternalServerError, &ResponseErr{Error: err.Error()})
		return
	}

	SendHTTPResponse(w, http.StatusOK, empty.Empty{})
}

func (db *dbServer) prepareAndCommitConfigTx(conf *config.Configurations) error {
	configTx, err := prepareConfigTx(conf)
	if err != nil {
		return errors.Wrap(err, "failed to prepare and commit a configuration transaction")
	}

	if err := db.submitTransaction(configTx); err != nil {
		return errors.Wrap(err, "error while committing configuration transaction")
	}
	return nil
}

func validateAndParseHeader(h *http.Header) (string, []byte, error) {
	userID := h.Get(constants.UserHeader)
	if userID == "" {
		return "", nil, errors.New(constants.UserHeader + " is not set in the http request header")
	}

	signature := h.Get(constants.SignatureHeader)
	if signature == "" {
		return "", nil, errors.New(constants.SignatureHeader + " is not set in the http request header")
	}
	signatureBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return "", nil, errors.New(constants.SignatureHeader + " is not encoded correctly")
	}

	return userID, signatureBytes, nil
}

type certsInGenesisConfig struct {
	nodeCert   []byte
	adminCert  []byte
	rootCACert []byte
}

func prepareConfigTx(conf *config.Configurations) (*types.ConfigTxEnvelope, error) {
	certs, err := readCerts(conf)
	if err != nil {
		return nil, err
	}

	clusterConfig := &types.ClusterConfig{
		Nodes: []*types.NodeConfig{
			{
				ID:          conf.Node.Identity.ID,
				Certificate: certs.nodeCert,
				Address:     conf.Node.Network.Address,
				Port:        conf.Node.Network.Port,
			},
		},
		Admins: []*types.Admin{
			{
				ID:          conf.Admin.ID,
				Certificate: certs.adminCert,
			},
		},
		RootCACertificate: certs.rootCACert,
	}

	return &types.ConfigTxEnvelope{
		Payload: &types.ConfigTx{
			TxID:      uuid.New().String(), // TODO: we need to change TxID to string
			NewConfig: clusterConfig,
		},
		// TODO: we can make the node itself sign the transaction
	}, nil
}

func readCerts(conf *config.Configurations) (*certsInGenesisConfig, error) {
	nodeCert, err := ioutil.ReadFile(conf.Node.Identity.CertificatePath)
	if err != nil {
		return nil, errors.Wrapf(err, "error while reading node certificate %s", conf.Node.Identity.CertificatePath)
	}
	nodePemCert, _ := pem.Decode(nodeCert)

	adminCert, err := ioutil.ReadFile(conf.Admin.CertificatePath)
	if err != nil {
		return nil, errors.Wrapf(err, "error while reading admin certificate %s", conf.Admin.CertificatePath)
	}
	adminPemCert, _ := pem.Decode(adminCert)

	rootCACert, err := ioutil.ReadFile(conf.RootCA.CertificatePath)
	if err != nil {
		return nil, errors.Wrapf(err, "error while reading rootCA certificate %s", conf.RootCA.CertificatePath)
	}
	rootCAPemCert, _ := pem.Decode(rootCACert)

	return &certsInGenesisConfig{
		nodeCert:   nodePemCert.Bytes,
		adminCert:  adminPemCert.Bytes,
		rootCACert: rootCAPemCert.Bytes,
	}, nil
}

func createLedgerDir(dir string) error {
	exist, err := fileops.Exists(dir)
	if err != nil {
		return err
	}
	if exist {
		return nil
	}

	return fileops.CreateDir(dir)
}
