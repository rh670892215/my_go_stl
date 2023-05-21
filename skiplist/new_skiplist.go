package skiplist

// 新建skipList node
func newSkipList[K any, V any](level int, key K, value V) *node[K, V] {
	switch level {
	case 1:
		n := struct {
			head  node[K, V]
			nexts [1]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 2:
		n := struct {
			head  node[K, V]
			nexts [2]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 3:
		n := struct {
			head  node[K, V]
			nexts [3]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 4:
		n := struct {
			head  node[K, V]
			nexts [4]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 5:
		n := struct {
			head  node[K, V]
			nexts [5]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 6:
		n := struct {
			head  node[K, V]
			nexts [6]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 7:
		n := struct {
			head  node[K, V]
			nexts [7]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 8:
		n := struct {
			head  node[K, V]
			nexts [8]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 9:
		n := struct {
			head  node[K, V]
			nexts [9]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 10:
		n := struct {
			head  node[K, V]
			nexts [10]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 11:
		n := struct {
			head  node[K, V]
			nexts [11]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 12:
		n := struct {
			head  node[K, V]
			nexts [12]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 13:
		n := struct {
			head  node[K, V]
			nexts [13]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 14:
		n := struct {
			head  node[K, V]
			nexts [14]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 15:
		n := struct {
			head  node[K, V]
			nexts [15]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 16:
		n := struct {
			head  node[K, V]
			nexts [16]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 17:
		n := struct {
			head  node[K, V]
			nexts [17]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 18:
		n := struct {
			head  node[K, V]
			nexts [18]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 19:
		n := struct {
			head  node[K, V]
			nexts [19]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 20:
		n := struct {
			head  node[K, V]
			nexts [20]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 21:
		n := struct {
			head  node[K, V]
			nexts [21]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 22:
		n := struct {
			head  node[K, V]
			nexts [22]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 23:
		n := struct {
			head  node[K, V]
			nexts [23]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 24:
		n := struct {
			head  node[K, V]
			nexts [24]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 25:
		n := struct {
			head  node[K, V]
			nexts [25]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 26:
		n := struct {
			head  node[K, V]
			nexts [26]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 27:
		n := struct {
			head  node[K, V]
			nexts [27]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 28:
		n := struct {
			head  node[K, V]
			nexts [28]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 29:
		n := struct {
			head  node[K, V]
			nexts [29]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 30:
		n := struct {
			head  node[K, V]
			nexts [30]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 31:
		n := struct {
			head  node[K, V]
			nexts [31]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 32:
		n := struct {
			head  node[K, V]
			nexts [32]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 33:
		n := struct {
			head  node[K, V]
			nexts [33]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 34:
		n := struct {
			head  node[K, V]
			nexts [34]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 35:
		n := struct {
			head  node[K, V]
			nexts [35]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 36:
		n := struct {
			head  node[K, V]
			nexts [36]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 37:
		n := struct {
			head  node[K, V]
			nexts [37]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 38:
		n := struct {
			head  node[K, V]
			nexts [38]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 39:
		n := struct {
			head  node[K, V]
			nexts [39]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	case 40:
		n := struct {
			head  node[K, V]
			nexts [40]*node[K, V]
		}{head: node[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
	}

	panic("should not reach here")
}
