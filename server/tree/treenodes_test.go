package tree
import "testing"

func TestCloningOfTree(t *testing.T) {
  	// TODO build a complete tree and see if clone works properly
	left := NewUserLeafNode("lefty", "leftiesPubKey", 1, 1, make([]byte))
	right := NewUserLeafNode("righty", "rightiesPubKey", 1, 1, make([]byte))
	root := NewRootNode(left, right, 0, []byte("prev"), 1, []byte("lh"),  []byte("rh"))
	clone := root.Clone(1, 2)
	if clone != nil {
		t.Log("Cloned")
	}
}
