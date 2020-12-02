package mtree

import "github.com/pkg/errors"

// Proof calculate intermediate hashes between leaf with given index and root (caller node)
func (n *Node) Proof(leafIndex int) ([][]byte, error) {
	path, err := n.findPath(leafIndex)
	if err != nil {
		return nil, errors.Wrapf(err, "")
	}
	return computeProofFromPath(path), nil
}

func (n *Node) findPath(leafIndex int) ([]*Node, error) {
	if n == nil {
		return nil, errors.Errorf("no paths exist in nil tree")
	}
	if n.Left() == nil {
		if n.leafRange.endIndex == n.leafRange.startIndex && n.leafRange.endIndex == leafIndex {
			return []*Node{n}, nil
		} else {
			return nil, errors.Errorf("node with index %d is not part of merkle tree (%d, %d)", leafIndex, n.leafRange.startIndex, n.leafRange.endIndex)
		}
	}
	if n.leafRange.startIndex > leafIndex || n.leafRange.endIndex < leafIndex {
		return nil, errors.Errorf("node with index %d is not part of merkle tree (%d, %d)", leafIndex, n.leafRange.startIndex, n.leafRange.endIndex)
	}

	var subPath []*Node
	var err error
	if n.Left().leafRange.endIndex >= leafIndex {
		subPath, err = n.Left().findPath(leafIndex)
	} else {
		subPath, err = n.Right().findPath(leafIndex)
	}
	if err != nil {
		return nil, err
	}
	return append(subPath, n), nil
}

func computeProofFromPath(path []*Node) [][]byte {
	if len(path) == 0 {
		return nil
	}
	proof := make([][]byte, 0)

	for _, n := range path[:len(path)-1] {
		siblingHash := n.Sibling().Hash()
		if siblingHash != nil {
			proof = append(proof, siblingHash)
		}
	}

	return proof
}