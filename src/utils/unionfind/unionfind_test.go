package unionfind

import (
	"fmt"
	"testing"
)

func TestUnionFind(t *testing.T) {
	union := New(5)
	union.Merge(1, 2)
	union.Merge(1, 3)
	union.Merge(0, 4)
	fmt.Println(union)
	if union.Find(3) != union.Find(2) {
		t.Error("Expected same got ", union.Find(2), "!=", union.Find(3))
	}
}
