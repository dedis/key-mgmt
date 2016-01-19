package tree

import "log"


type TreeNodeI interface {
	Left() *TreeNodeI
	Right() *TreeNodeI
	Parent() *TreeNodeI
	Level() int
	Name()  string
	Clone(parent *TreeNodeI, curEpoch, nextEpoch int) *TreeNodeI
}


// TreeNode Represents an generic tree node in the CONIKS binary Merkle prefix tree.
type TreeNode struct {
	// XXX for simplicity everything is public <- change that soon
	left   *TreeNodeI
	right  *TreeNodeI

	parent *TreeNodeI
	level  int

	name   string
}

func (tn *TreeNode) Left() *TreeNodeI {
	return tn.left
}

func (tn *TreeNode) Right() *TreeNodeI {
	return tn.right
}

func (tn *TreeNode) Parent() *TreeNodeI {
	return tn.parent
}

func (tn *TreeNode) Level() int {
	return tn.level
}

func (tn *TreeNode) Name() string {
	return tn.name
}

func (tn *TreeNode) Clone(parent *TreeNodeI, curEpoch, nextEpoch int) *TreeNodeI {
	panic("Shouldn't be called")
	return nil
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

func (tn *UserLeafNode) Left() *TreeNodeI {
	return tn.left
}

func (tn *UserLeafNode) Right() *TreeNodeI {
	return tn.right
}

func (tn *UserLeafNode) Parent() *TreeNodeI {
	return tn.parent
}

func (tn *UserLeafNode) Level() int {
	return tn.level
}

func (tn *UserLeafNode) Name() string {
	return tn.name
}



// NewUserLeafNode create a UserLeafNode from the given arguments
func NewUserLeafNode(username string, pub string, epoch int, lvl int, index []byte) *UserLeafNode {
	return &UserLeafNode{
		LeafNode: &LeafNode{
			left: nil,
			right: nil,
			parent: nil,
			level: lvl,
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


// Clone duplicates the given user leaf node from the current epoch for the next epoch with the
// given parent tree node.
//
// This function is called as part of the CONIKS Merkle tree
// rebuilding process at the beginning of every epoch.
func (uln *UserLeafNode) Clone(parent *TreeNodeI, curEpoch int, nextEpoch int) *TreeNodeI {
	log.Println("UserLeafNode Clone() called")
	// FIXME epochs aren't even used (compare UserLeafNode.java)
	var cloneN TreeNodeI
	cloneN = NewUserLeafNode(uln.Username, uln.PubKey, uln.EpochAdded, uln.level, uln.Index)
	//cloneN.parent = parent

	return &cloneN
}

type RootNode struct {
	*InteriorNode
	Prev  []byte
	Epoch int
}

// NewRootNode constructs a root node specified  with left and right subtrees,
// the hash of the previous epoch's tree root {@code prev},
// the level in tree {@code lvl}, and the epoch {@code ep} for which this root is valid.
func NewRootNode(left TreeNodeI, right TreeNodeI, lvl int, prev []byte, ep int, lh ,rh []byte) *RootNode {
	return &RootNode{
		InteriorNode: &InteriorNode{
			TreeNode: &TreeNode{
				left: &left,
				right: &right,
				level: lvl,
			},
			LeftHash:  lh,
			RightHash: rh,
			hasLeaf: false, // XXX refactor, compare RootNode.java
		},
		Prev: prev,
		Epoch: ep,
	}
}



// TODO PaddingLeafNode (currently not even used in the reference implementation)
