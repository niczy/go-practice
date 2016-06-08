package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int               { return len(h) }
func (h IntHeap) Less(i int, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i int, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = (*h)[0 : l-1]
	return x
}

func smallestKElement(nums []int, k int) []int {
	h := &IntHeap{}
	heap.Init(h)
	//fmt.Printf("nums %v\n", nums)
	for _, num := range nums {
		heap.Push(h, num)
		//fmt.Printf("heap %v\n", h)
	}
	ret := make([]int, 0)
	for i := 0; i < k; i++ {
		ret = append(ret, heap.Pop(h).(int))
	}
	return ret
}

func main() {
	fmt.Printf("%v\n", smallestKElement([]int{1, 4, 3, 5, 9, -1}, 4))
}
