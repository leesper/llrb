package llrb

import (
  "strings"
  "testing"
)

func TestStringKeyRBTreeInsertion(t *testing.T) {
    wordForWord := NewLowerCaseRBTree()
    for _, word := range []string{"one", "Two", "THREE", "four", "Five"} {
        wordForWord.Insert(word, word)
    }
    var words []string
    wordForWord.Do(func(_, value interface{}) {
        words = append(words, value.(string))
    })
    actual, expected := strings.Join(words, ""), "FivefouroneTHREETwo"
    if actual != expected {
        t.Errorf("%q != %q", actual, expected)
    }
}

func TestIntKeyRBTreeFind(t *testing.T) {
    intMap := NewIntRBTree()
    for _, number := range []int{9, 1, 8, 2, 7, 3, 6, 4, 5, 0} {
        intMap.Insert(number, number*10)
    }
    for _, number := range []int{0, 1, 5, 8, 9} {
        value, found := intMap.Find(number)
        if !found {
            t.Errorf("failed to find %d", number)
        }
        actual, expected := value.(int), number*10
        if actual != expected {
            t.Errorf("value is %d should be %d", actual, expected)
        }
    }
    for _, number := range []int{-1, -21, 10, 11, 148} {
        _, found := intMap.Find(number)
        if found {
            t.Errorf("should not have found %d", number)
        }
    }
}

func TestIntKeyRBTreeDelete(t *testing.T) {
    intMap := NewIntRBTree()
    for _, number := range []int{9, 1, 8, 2, 7, 3, 6, 4, 5, 0} {
        intMap.Insert(number, number*10)
    }
    if intMap.Len() != 10 {
        t.Errorf("map len %d should be 10", intMap.Len())
    }
    length := 9
    for i, number := range []int{0, 1, 5, 8, 9} {
        if deleted := intMap.Delete(number); !deleted {
            t.Errorf("failed to delete %d", number)
        }
        if intMap.Len() != length-i {
            t.Errorf("map len %d should be %d", intMap.Len(), length-i)
        }
    }
    for _, number := range []int{-1, -21, 10, 11, 148} {
        if deleted := intMap.Delete(number); deleted {
            t.Errorf("should not have deleted nonexistent %d", number)
        }
    }
    if intMap.Len() != 5 {
        t.Errorf("map len %d should be 5", intMap.Len())
    }
}

func TestPassing(t *testing.T) {
    intRBTree := NewIntRBTree()
    intRBTree.Insert(7, 7)
    passRBTree(intRBTree, t)
}

func passRBTree(r *RBTree, t *testing.T) {
    for _, number := range []int{9, 3, 6, 4, 5, 0} {
        r.Insert(number, number)
    }
    if r.Len() != 7 {
        t.Errorf("should have %d items", 7)
    }
}

// Thanks to Russ Cox for improving these benchmarks
func BenchmarkRBTreeFindSuccess(b *testing.B) {
    b.StopTimer() // Don't time creation and population
    intRBTree := NewIntRBTree()
    for i := 0; i < 1e6; i++ {
        intRBTree.Insert(i, i)
    }
    b.StartTimer() // Time the Find() method succeeding
    for i := 0; i < b.N; i++ {
        intRBTree.Find(i % 1e6)
    }
}

func BenchmarkRBTreeFindFailure(b *testing.B) {
    b.StopTimer() // Don't time creation and population
    intRBTree := NewIntRBTree()
    for i := 0; i < 1e6; i++ {
        intRBTree.Insert(2*i, i)
    }
    b.StartTimer() // Time the Find() method failing
    for i := 0; i < b.N; i++ {
        intRBTree.Find(2*(i%1e6) + 1)
    }
}
