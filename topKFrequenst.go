package main

import (
	"container/heap"
	"fmt"
)

type Node struct {
	Val   int
	Times int
}

type NodeHeap []Node

func (h NodeHeap) Len() int               { return len(h) }
func (h NodeHeap) Less(i int, j int) bool { return h[i].Times > h[j].Times }
func (h NodeHeap) Swap(i int, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = (*h)[0 : l-1]
	return x
}

func topKFrequentElement(nums []int, k int) []int {
	h := &NodeHeap{}
	heap.Init(h)
	//fmt.Printf("nums %v\n", nums)
	m := make(map[int]int)
	for _, num := range nums {
		val, ok := m[num]
		if ok {
			m[num] = val + 1
		} else {
			m[num] = 0
		}
		// heap.Push(h, num)
		//fmt.Printf("heap %v\n", h)
	}
	for k, v := range m {
		heap.Push(h, Node{
			Val:   k,
			Times: v,
		})
	}
	ret := make([]int, 0)
	for i := 0; i < k; i++ {
		ret = append(ret, heap.Pop(h).(Node).Val)
	}
	return ret
}

func main() {
	fmt.Printf("%v\n", topKFrequentElement([]int{1, 1, 1, 2, 2, -1}, 2))
}
