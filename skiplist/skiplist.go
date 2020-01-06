package skiplist

import (
	"math"
	"math/rand"
)

type (
	SkipList struct {
		head      *element
		cacheList []*element
		maxLevel  int
	}

	element struct {
		next []*element
		key  int
		val  interface{}
	}
)

func NewSkipList(maxLevel int) *SkipList {
	return &SkipList{
		head: &element{
			next: make([]*element, maxLevel),
			key:  math.MinInt32,
			val:  nil,
		},
		cacheList: make([]*element, maxLevel),
		maxLevel:  maxLevel,
	}
}

func (sl *SkipList) Set(key int, val interface{}) {
	var (
		tem *element
		pre = sl.head
	)

	for i := sl.maxLevel - 1; i >= 0; i-- {
		tem = pre.next[i]
		for tem != nil && tem.key < key {
			pre = tem
			tem = tem.next[i]
		}
		sl.cacheList[i] = pre
	}

	if sl.cacheList[0].next[0].key == key {
		sl.cacheList[0].next[0].val = val
		return
	}

	insertData := &element{
		next: make([]*element, sl.randomLevel()),
		key:  key,
		val:  val,
	}

	for i := 0; i < len(insertData.next); i++ {
		insertData.next[i] = sl.cacheList[0].next[i].next[i]
		sl.cacheList[0].next[i] = insertData
	}
}

func (sl *SkipList) Get(key int) (interface{}, bool) {
	var (
		pre = sl.head
		tem *element
	)

	for i := sl.maxLevel - 1; i >= 0; i-- {
		tem = pre.next[i]
		for tem != nil && tem.key < key {
			pre = tem
			tem = tem.next[i]
		}
	}

	if tem != nil && tem.key == key {
		return tem.val, true
	}

	return nil, false
}

func (sl *SkipList) Del(key int) (interface{}, bool) {
	var (
		pre = sl.head
		tem *element
	)

	for i := sl.maxLevel - 1; i >= 0; i-- {
		tem = pre.next[i]
		for tem != nil && tem.key < key {
			pre = tem
			tem = tem.next[i]
		}
		sl.cacheList[i] = pre
	}

	if sl.cacheList[0].next[0].key != key {
		return nil, false
	}

	val := sl.cacheList[0].next[0].val

	for i := 0; i < len(sl.cacheList[0].next); i++ {
		sl.cacheList[i].next[i] = sl.cacheList[i].next[i].next[i]
	}

	return val, true
}

func (sl *SkipList) randomLevel() int {
	level := 1

	for {
		if rand.Intn(1) == 0 {
			level++
		} else {
			return level
		}

		if level >= sl.maxLevel {
			return sl.maxLevel
		}

	}
}
