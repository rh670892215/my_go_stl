package skiplist

import "testing"

const (
	benchInitSize  = 1000000
	benchBatchSize = 10
)

func newSkipListN(n int) *SkipList[int, int] {
	sl := NewSkipList[int, int]()
	sl.random.Seed(0)
	for i := 0; i < n; i++ {
		sl.Insert(i, i)
	}
	return sl
}

func BenchmarkSkipList_Insert(b *testing.B) {
	start := benchInitSize
	sl := newSkipListN(start)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := 0; i < benchBatchSize; i++ {
			sl.Insert(start+i, i)
		}
		start += benchBatchSize
	}
}
