package main

import (
	"fmt"
)

func solve(n int, edges [][2]int) []int {
	g := make([][]int, n)
	/*
		for i := 0; i < n; i++ {
			g[i] = make([]int, 0)
		}
	*/
	for i := 0; i < len(edges); i++ {
		p := edges[i][0]
		q := edges[i][1]
		g[p] = append(g[p], q)
		g[q] = append(g[q], p)
	}
	removed := make(map[int]bool)
	for {
		ret := make([]int, 0)
		for i := 0; i < n; i++ {
			if !removed[i] {
				ret = append(ret, i)
			}
		}
		if len(ret) <= 2 {
			return ret
		}
		leaf := make([]int, 0)
		for i := 0; i < n; i++ {
			if len(g[i]) == 1 {
				leaf = append(leaf, i)
				removed[i] = true
			}
		}
		fmt.Printf("leaves %v\n", leaf)
		for _, r := range leaf {
			ne := g[r][0]
			g[r] = g[r][0:0]
			for idx, j := range g[ne] {
				if j == r {
					g[ne] = append(g[ne][0:idx], g[ne][idx+1:]...)
					break
				}
			}
		}

	}

}

func main() {
	fmt.Printf("%v\n", solve(2, [][2]int{{0, 1}}))
	fmt.Printf("%v\n", solve(1, [][2]int{}))
	fmt.Printf("%v\n", solve(6, [][2]int{{0, 3}, {1, 3}, {2, 3}, {4, 3}, {5, 4}}))
	fmt.Printf("%v\n", solve(7, [][2]int{{5, 6}, {0, 3}, {1, 3}, {2, 3}, {4, 3}, {5, 4}}))
}
