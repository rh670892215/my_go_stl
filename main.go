package main

import (
	"fmt"
	"my_go_stl/skiplist"

	"github.com/chen3feng/stl4go"
)

func main() {
	sl := skiplist.NewSkipList[int, int]()
	newSl := stl4go.NewSkipList[int, int]()
	for i := 0; i < 10; i++ {
		if i == 8 {
			continue
		}
		sl.Insert(i, i)
		newSl.Insert(i, i)
	}
	for i := 0; i < 10; i++ {
		res := sl.Get(i)
		fmt.Println(*res)
	}
	fmt.Println("here")
}
