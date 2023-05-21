package skiplist

import (
	"math/bits"
	"math/rand"
	"my_go_stl/types"
	"time"
)

const (
	SkipListMaxLevel = 40
)

// 跳表节点
type node[K, V any] struct {
	key   K
	value V
	next  []*node[K, V]
}

// SkipList 跳表
type SkipList[K, V any] struct {
	head         node[K, V]
	level        int
	len          int
	random       *rand.Rand
	skipListImpl skipListImpl[K, V]
	prevsCache   []*node[K, V]
}

// NewSkipList 新建跳表
func NewSkipList[K types.Order, V any]() *SkipList[K, V] {
	sl := skipListOrder[K, V]{}
	sl.init()
	sl.skipListImpl = (skipListImpl[K, V])(&sl)
	return &sl.SkipList
}

// 一些对外暴露的接口

func (sl *SkipList[K, V]) Len() int {
	return sl.len
}

func (sl *SkipList[K, V]) IsEmpty() bool {
	return sl.len == 0
}

func (sl *SkipList[K, V]) Clear() {
	for i, _ := range sl.head.next {
		sl.head.next[i] = nil
	}
	sl.len = 0
	sl.level = 1
}

// Get 查找节点
func (sl *SkipList[K, V]) Get(key K) *V {
	targetNode := sl.skipListImpl.findNode(key)
	if targetNode == nil {
		return nil
	}
	return &targetNode.value

}

// Insert 插入节点
func (sl *SkipList[K, V]) Insert(key K, value V) {
	targetNode, prevNodes := sl.skipListImpl.findInsertPoint(key)
	if targetNode != nil {
		targetNode.value = value
		return
	}

	targetLevel := sl.randomLevel()
	targetNode = newSkipList[K, V](targetLevel, key, value)

	for i := 0; i < types.Min(targetLevel, sl.level); i++ {
		targetNode.next[i] = prevNodes[i].next[i]
		prevNodes[i].next[i] = targetNode
	}

	if targetLevel > sl.level {
		for i := sl.level; i < targetLevel; i++ {
			sl.head.next[i] = targetNode
		}
		sl.level = targetLevel
	}
	sl.len++
}

// Remove 移除节点
func (sl *SkipList[K, V]) Remove(key K) bool {
	targetNode, targetNodePrevs := sl.skipListImpl.findRemovePoint(key)
	if targetNode == nil {
		return false
	}
	for i, cur := range targetNode.next {
		targetNodePrevs[i].next[i] = cur
	}
	//调整level,这里要用for循环逐层清理,不要用if,否则会只清理顶层
	for sl.level > 1 && sl.head.next[sl.level-1] == nil {
		sl.level--
	}
	sl.len--
	return true
}

// 跳表相关内部方法的实现接口
type skipListImpl[K, V any] interface {
	findNode(key K) *node[K, V]
	findInsertPoint(key K) (*node[K, V], []*node[K, V])
	findRemovePoint(key K) (*node[K, V], []*node[K, V])
}

func (sl *SkipList[K, V]) init() {
	sl.level = 1
	sl.len = 0
	sl.random = rand.New(rand.NewSource(time.Now().Unix()))
	sl.head.next = make([]*node[K, V], SkipListMaxLevel)
	sl.prevsCache = make([]*node[K, V], SkipListMaxLevel)
}

func (sl *SkipList[K, V]) randomLevel() int {
	//方案一
	// total := uint64(1)<<uint64(SkipListMaxLevel) - 1 // 2^n - 1
	// k := sl.random.Uint64() % total
	// levelDivide := uint64(1) << (uint64(SkipListMaxLevel) - 1)
	// level := 1
	// for total -= levelDivide; total > k; level++ {
	// 	levelDivide >>= 1
	// 	total -= levelDivide
	// }
	//return level

	//方案二
	total := uint64(1)<<uint64(SkipListMaxLevel) - 1 // 2^n-1
	k := sl.random.Uint64() % total
	level := SkipListMaxLevel - bits.Len64(k) + 1
	for level > 2 && 1<<(level-2) > sl.len {
		level--
	}

	return level
}

type skipListOrder[K types.Order, V any] struct {
	SkipList[K, V]
}

func (slo *skipListOrder[K, V]) findNode(key K) *node[K, V] {
	return slo.doFindNode(key, true)
}

// 查找指定key的节点是否存在
func (slo *skipListOrder[K, V]) doFindNode(key K, eq bool) *node[K, V] {
	prev := &slo.head

	// 从上往下找,逐层级遍历
	for i := slo.level - 1; i >= 0; i-- {
		for cur := prev.next[i]; cur != nil; cur = cur.next[i] {
			if cur.key == key {
				return cur
			}
			// 该节点的key已经大于目标节点,本层不可能存在目标节点了,继续去下一层寻找
			if cur.key > key {
				break
			}
			prev = cur
		}
	}
	if eq {
		return nil
	}
	return prev.next[0]
}

// 查找待插入节点的插入位置,如果节点已存在则返回该节点,如果节点不存在则返回待插入位置的前置指针
func (slo *skipListOrder[K, V]) findInsertPoint(key K) (*node[K, V], []*node[K, V]) {
	//prevs := make([]*node[K, V], slo.level)
	prevs := slo.prevsCache[0:slo.level]
	prev := &slo.head

	for i := slo.level - 1; i >= 0; i-- {
		for next := prev.next[i]; next != nil; next = next.next[i] {
			if next.key == key {
				return next, nil
			}
			if next.key > key {
				break
			}
			prev = next
		}
		prevs[i] = prev
	}

	return nil, prevs
}

// 查找待删除节点的位置
func (slo *skipListOrder[K, V]) findRemovePoint(key K) (*node[K, V], []*node[K, V]) {
	prevs := slo.findPrePoint(key)
	targetNode := prevs[0].next[0]

	if targetNode == nil || targetNode.key != key {
		return nil, nil
	}
	return targetNode, prevs
}

// 查找指定节点的前置节点集合
func (slo *skipListOrder[K, V]) findPrePoint(key K) []*node[K, V] {
	//prevs := make([]*node[K, V], slo.level)
	prevs := slo.prevsCache[0:slo.level]
	prev := &slo.head

	for i := slo.level - 1; i >= 0; i-- {
		for next := prev.next[i]; next != nil; next = next.next[i] {
			if next.key >= key {
				break
			}
			prev = next
		}
		prevs[i] = prev
	}

	return prevs
}
