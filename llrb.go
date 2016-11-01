package llrb

import (
  "fmt"
)

type rbnode struct {
  key, value interface{}
  red bool
  left, right *rbnode
}

func newRBNode(k, v interface{}) *rbnode {
  return &rbnode{
    key: k,
    value: v,
  }
}

func isRed(node *rbnode) bool {
  return node != nil && node.red
}

func rotateLeft(node *rbnode) *rbnode {
  rightChild := node.right
  node.right = rightChild.left
  rightChild.left = node
  rightChild.red = node.red
  node.red = true
  return rightChild
}

func rotateRight(node *rbnode) *rbnode {
  leftChild := node.left
  node.left = leftChild.right
  leftChild.right = node
  leftChild.red = node.red
  node.red = true
  return leftChild
}

func colorFlip(node *rbnode) *rbnode {
  node.red = !node.red

  if node.left != nil {
    node.left.red = !node.left.red
  }

  if node.right != nil {
    node.right.red = !node.right.red
  }

  return node
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

func printInorderTree(node *rbnode, l int) {
  l++
  if node == nil {
    printBlank(l)
    fmt.Println("Nil")
  } else {
    printBlank(l)
    fmt.Printf("(%v%t\n", node.key, node.red)
    printInorderTree(node.left, l)
    printInorderTree(node.right, l)
    printBlank(l)
    fmt.Println(")")
  }
}

func printBlank(l int) {
  for i := 0; i < l; i++ {
    fmt.Printf(" ")
  }
}
