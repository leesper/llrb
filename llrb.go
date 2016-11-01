package llrb

type color int8

const (
  RED color = 1
  BLACK color = 0
)

type rbnode struct {
  key, value interface{}
  color color
  parent, left, right *rbnode
}

func (node *rbnode) isRed() bool {
  if node == nil {
    return false
  }
  return node.color == RED
}

// root's right child point to the root of RBTree
type RBTree struct {
  root *rbnode
  length int
}

func NewRBTree() RBTree {
  return RBTree{
    root: new(rbnode),
  }
}
