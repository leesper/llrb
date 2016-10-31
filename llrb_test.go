package llrb

import (
  "testing"
)

func TestNewRBTree(t *testing.T) {
  rbTree := NewRBTree()
  if rbTree.root == nil {
    t.Errorf("rbTree is nil, want non-nil")
  }
  if rbTree.root.right != nil {
    t.Errorf("rbTree.root.right is non-nil, want nil")
  }
}
