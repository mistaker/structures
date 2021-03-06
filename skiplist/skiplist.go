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
		next *element
		prev = sl.head
	)

	for i := sl.maxLevel - 1; i >= 0; i-- {
		next = prev.next[i]
		for next != nil && next.key < key {
			prev = next
			next = next.next[i]
		}
		sl.cacheList[i] = prev
	}

	if elm := sl.cacheList[0].next[0]; elm != nil && elm.key == key {
		sl.cacheList[0].val = val
		return
	}

	insertData := &element{
		next: make([]*element, sl.randomLevel()),
		key:  key,
		val:  val,
	}

	for i := 0; i < len(insertData.next); i++ {
		insertData.next[i] = sl.cacheList[i].next[i]
		sl.cacheList[i].next[i] = insertData
	}
}

func (sl *SkipList) Get(key int) (interface{}, bool) {
	var (
		prev = sl.head
		next *element
	)

	for i := sl.maxLevel - 1; i >= 0; i-- {
		next = prev.next[i]
		for next != nil && next.key < key {
			prev = next
			next = next.next[i]
		}
	}

	if next != nil && next.key == key {
		return next.val, true
	}

	return nil, false
}

func (sl *SkipList) Del(key int) (interface{}, bool) {
	var (
		prev = sl.head
		next *element
	)

	for i := sl.maxLevel - 1; i >= 0; i-- {
		next = prev.next[i]
		for next != nil && next.key < key {
			prev = next
			next = next.next[i]
		}
		sl.cacheList[i] = prev
	}

	if elm := sl.cacheList[0].next[0]; elm == nil || elm.key != key {
		return nil, false
	}

	val := sl.cacheList[0].val

	for i := 0; i < len(sl.cacheList[0].next[0].next); i++ {
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
