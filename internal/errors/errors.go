// Copyright IBM Corp. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
package errors

import "fmt"

type NotFoundErr struct {
	Message string
}

func (e *NotFoundErr) Error() string {
	return e.Message
}

type PermissionErr struct {
	ErrMsg string
}

func (e *PermissionErr) Error() string {
	return e.ErrMsg
}

type TimeoutErr struct {
	ErrMsg string
}

func (t *TimeoutErr) Error() string {
	return t.ErrMsg
}

// DuplicateTxIDError is an error to denote that a transaction
// has a duplicate txID
type DuplicateTxIDError struct {
	TxID string
}

func (d *DuplicateTxIDError) Error() string {
	return "the transaction has a duplicate txID [" + d.TxID + "]"
}

// ClosedError is used when a blocking operation aborted because a component closed,
// or when an operation is performed on a component that is already closed.
type ClosedError struct {
	ErrMsg string
}

func (c *ClosedError) Error() string {
	return c.ErrMsg
}

// NotLeaderError is an error that denotes that the current node is not the cluster leader.
// The error carries the identity of the leader if it is known (>0), or 0 if it is not.
type NotLeaderError struct {
	LeaderID       uint64 // RaftID of the leader
	LeaderHostPort string // The leader's node host:port for client request redirect
}

func (n *NotLeaderError) Error() string {
	return fmt.Sprintf("not a leader, leader is RaftID: %d, with HostPort: %s", n.LeaderID, n.LeaderHostPort)
}

func (n *NotLeaderError) GetLeaderID() uint64 {
	return n.LeaderID
}

func (n *NotLeaderError) GetLeaderHostPort() string {
	return n.LeaderHostPort
}

// BadRequestError is used for errors that should be translated to a bad request, for example as an illegal TxId.
type BadRequestError struct {
	ErrMsg string
}

func (c *BadRequestError) Error() string {
	return c.ErrMsg
}