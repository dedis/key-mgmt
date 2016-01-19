package tree
import "testing"

func TestCloningOfTree(t *testing.T) {
  	// TODO build a complete tree and see if clone works properly
	left := NewUserLeafNode("lefty", "leftiesPubKey", 1, 1, []byte("TestIndex"))
	right := NewUserLeafNode("righty", "rightiesPubKey", 1, 1, []byte("TestIndex"))
	root := NewRootNode(left, right, 0, []byte("prev"), 1, []byte("lh"),  []byte("rh"))
	clone := root.Clone(nil, 1, 2)
	if clone != nil {
		t.Log("Cloned")
	}
}
