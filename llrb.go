package llrb

import (
  "fmt"
  "strings"
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

func moveRedLeft(node *rbnode) *rbnode {
  node = colorFlip(node)
  if isRed(node.right.left) {
    node.right = rotateRight(node.right)
    node = rotateLeft(node)
    node = colorFlip(node)
  }
  return node
}

func moveRedRight(node *rbnode) *rbnode {
  node = colorFlip(node)
  if node.left != nil && isRed(node.left.left) {
    node = rotateRight(node)
    node = colorFlip(node)
  }
  return node
}

func deleteMax(node *rbnode) *rbnode {
  if isRed(node.left) {
    node = rotateRight(node)
  }
  if node.right == nil {
    return nil
  }
  if !isRed(node.right) && !isRed(node.right.left) {
    node = moveRedRight(node)
  }
  node.right = deleteMax(node.right)
  return fixUp(node)
}

func deleteMin(node *rbnode) *rbnode {
  if node.left == nil {
    return nil
  }
  if !isRed(node.left) && !isRed(node.left.left) {
    node = moveRedLeft(node)
  }
  node.left = deleteMin(node.left)
  return fixUp(node)
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

func fixUp(node *rbnode) *rbnode {
  if isRed(node.right) {
    node = rotateLeft(node)
  }
  if isRed(node.left) && isRed(node.left.left) {
    node = rotateRight(node)
  }
  if isRed(node.left) && isRed(node.right) {
    node = colorFlip(node)
  }
  return node
}

// root's right child point to the root of RBTree
type RBTree struct {
  root *rbnode
  length int
  less func(a, b interface{}) bool
}

func NewRBTree(less func(a, b interface{}) bool) *RBTree {
  return &RBTree{
    less: less,
  }
}

func NewLowerCaseRBTree() *RBTree {
  return &RBTree{less: func(a, b interface{}) bool {
      return strings.ToLower(a.(string)) < strings.ToLower(b.(string))
  }}
}

func NewIntRBTree() *RBTree {
  return &RBTree{less: func(a, b interface{}) bool {
    return a.(int) < b.(int)
  }}
}

func NewFloat64RBTree() *RBTree {
  return &RBTree{less: func(a, b interface{}) bool {
    return a.(float64) < b.(float64)
  }}
}

func (r *RBTree) Insert(k, v interface{}) bool {
  ok := false
  r.root, ok = r.insert(r.root, k, v)
  r.root.red = false
  if ok {
    r.length++
  }
  return ok
}

func (r *RBTree) Find(k interface{}) (interface{}, bool) {
  root := r.root
  for root != nil {
    if r.less(k, root.key) {
      root = root.left
    } else if r.less(root.key, k) {
      root = root.right
    } else {
      return root.value, true
    }
  }
  return nil, false
}

func (r *RBTree) insert(node *rbnode, k, v interface{}) (*rbnode, bool) {
  ok := false
  if node == nil {
    return &rbnode{key: k, value: v, red: true}, true
  }
  if r.less(k, node.key) {
    node.left, ok = r.insert(node.left, k, v)
  } else if r.less(node.key, k) {
    node.right, ok = r.insert(node.right, k, v)
  } else {
    node.value = v
  }
  return fixUp(node), ok
}

func (r *RBTree) Do(function func(interface{}, interface{})) {
  do(r.root, function)
}

func (r *RBTree) Len() int {
  return r.length
}

func do(node *rbnode, function func(interface{}, interface{})) {
  if node != nil {
    do(node.left, function)
    function(node.key, node.value)
    do(node.right, function)
  }
}

func (r *RBTree) Delete(k interface{}) bool {
  ok := false
  if r.root != nil {
    if r.root, ok = r.delete(r.root, k); r.root != nil {
      r.root.red = false
    }
  }
  if ok {
    r.length--
  }
  return ok
}

func (r *RBTree) delete(node *rbnode, k interface{}) (*rbnode, bool) {
  ok := false
  if r.less(k, node.key) {
    if node.left != nil {
      if !isRed(node.left) && !isRed(node.left.left) {
        node = moveRedLeft(node)
      }
      node.left, ok = r.delete(node.left, k)
    }
  } else {
    if isRed(node.left) {
      node = rotateRight(node)
    }
    if !r.less(k, node.key) && !r.less(node.key, k) && node.right == nil {
      return nil, true
    }
    if node.right != nil {
      if !isRed(node.right) && !isRed(node.right.left) {
        node = moveRedRight(node)
      }
      if !r.less(k, node.key) && !r.less(node.key, k) {
        smallest := min(node.right)
        node.key = smallest.key
        node.value = smallest.value
        node.right = deleteMin(node.right)
        ok = true
      } else {
        node.right, ok = r.delete(node.right, k)
      }
    }
  }
  return fixUp(node), ok
}

func min(node *rbnode) *rbnode {
  for node.left != nil {
    node = node.left
  }
  return node
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
