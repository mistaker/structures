package skiplist

import (
	"fmt"
	"testing"
)

func TestSkipList_Dels(t *testing.T) {
	sl := NewSkipList(5)
	sl.Set(1, 1)

	fmt.Println(sl.Get(1))

}
