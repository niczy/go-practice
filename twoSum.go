package main

import (
	"fmt"
	"sort"
)

type Node struct {
	Val int
	Idx int
}

type NodeArr []Node

func (this NodeArr) Len() int               { return len(this) }
func (this NodeArr) Less(i int, j int) bool { return this[i].Val < this[j].Val }
func (this NodeArr) Swap(i int, j int)      { this[i], this[j] = this[j], this[i] }

func twoSum(nums []int, target int) []int {
	n := NodeArr{}
	for idx, num := range nums {
		n = append(n, Node{
			Val: num,
			Idx: idx,
		})
	}
	sort.Sort(&n)
	i := 0
	j := len(nums) - 1
	for n[i].Val+n[j].Val != target {
		if n[i].Val+n[j].Val < target {
			i += 1
		} else {
			j -= 1
		}
	}
	return []int{n[i].Idx, n[j].Idx}
}

func main() {
	fmt.Printf("%v\n", twoSum([]int{2, 7, 11, 15}, 9))
}
