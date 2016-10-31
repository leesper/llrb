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

// root's right child point to the root of RBTree
type RBTree struct {
  root *rbnode
  length int
}
