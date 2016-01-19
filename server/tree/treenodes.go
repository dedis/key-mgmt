package tree

import "log"

// TreeNode Represents an generic tree node in the CONIKS binary Merkle prefix tree.
type TreeNode struct {
	// XXX for simplicity everything is public <- change that soon
	Left   *TreeNode
	Right  *TreeNode

	Parent *TreeNode
	Level  int

	Name   string
}

func (tn *TreeNode) Clone(parent *TreeNode, curEpoch int, nextEpoch int) *UserLeafNode {
	panic("Unsupported method call")
}

// LeafNode does not add any functionality, just use an alias
type LeafNode TreeNode

type InteriorNode struct {
	*TreeNode
	LeftHash  []byte
	RightHash []byte
	hasLeaf   bool
}

// Represents a leaf node containing a user's data binding interior node
// in the CONIKS binary Merkle prefix tree.
type UserLeafNode struct {
	*LeafNode

	Username               string
	PubKey                 string
	EpochAdded             int  // XXX long?

	AllowUnsignedKeychange bool // FIXME this is always true in the reference implementation ...
	AllowPublicLookup      bool

	Index                  []byte
	Signature              []byte

}

// NewUserLeafNode create a UserLeafNode from the given arguments
func NewUserLeafNode(username string, pub string, epoch int, lvl int, index []byte) *UserLeafNode {
	return &UserLeafNode{
		LeafNode: &LeafNode{
			Left: nil,
			Right: nil,
			Parent: nil,
			Level: lvl,
		},
		Username: username,
		PubKey: pub,
		EpochAdded: epoch,
		AllowPublicLookup: true, // this is the default in the reference implementation
		AllowUnsignedKeychange:true,
		Index: index,
		Signature: make([]byte, 256), // reference implementation says "dummy array"
	}
}

type Cloneable interface {
	Clone(*TreeNode, int, int) *TreeNode
}

// Clone duplicates the given user leaf node from the current epoch for the next epoch with the
// given parent tree node.
//
// This function is called as part of the CONIKS Merkle tree
// rebuilding process at the beginning of every epoch.
func (uln *UserLeafNode) Clone(parent *TreeNode, curEpoch int, nextEpoch int) *UserLeafNode {
	log.Println("UserLeafNode Clone() called")
	// FIXME epochs aren't even used (compare UserLeafNode.java)
	cloneN := NewUserLeafNode(uln.Username, uln.PubKey, uln.EpochAdded, uln.Level, uln.Index)
	cloneN.Parent = parent

	return cloneN
}

type RootNode struct {
	*InteriorNode
	Prev  []byte
	Epoch int
}

// NewRootNode constructs a root node specified  with left and right subtrees,
// the hash of the previous epoch's tree root {@code prev},
// the level in tree {@code lvl}, and the epoch {@code ep} for which this root is valid.
func NewRootNode(left TreeNode, right TreeNode, lvl int, prev []byte, ep int, lh ,rh []byte) *RootNode {
	return &RootNode{
		InteriorNode: &InteriorNode{
			TreeNode: &TreeNode{
				Left: &left,
				Right: &right,
				Level: lvl,
			},
			LeftHash:  lh,
			RightHash: rh,
			hasLeaf: false, // XXX refactor, compare RootNode.java
		},
		Prev: prev,
		Epoch: ep,
	}
}

func (rn *RootNode) Clone(curEpoch, nextEpoch int) *RootNode {
	log.Println("RootNode Clone() called")
	cloneN := NewRootNode(nil, nil, rn.Level, nil, -1, rn.LeftHash, rn.RightHash)
	if rn.Left != nil { // XXX I suspect this won't work?
		cloneN.Left = rn.Left.Clone(cloneN, curEpoch, nextEpoch)
	}
	if rn.Right != nil {
		cloneN.Right = rn.Right.Clone(cloneN, curEpoch, nextEpoch)
	}
	return cloneN
}


// TODO PaddingLeafNode (currently not even used in the reference implementation)
