package main

import (
	"fmt"
	"math"
)

func integerBreak(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	var p3 int
	var rest int
	if n%3 == 0 {
		p3 = n / 3
		rest = 1
	} else if n%3 == 1 {
		p3 = n/3 - 1
		rest = 4
	} else if n%3 == 2 {
		p3 = n / 3
		rest = 2
	}
	return int(math.Pow(float64(3), float64(p3))) * rest
}

func main() {
	m := make([]int, 100)
	ans := make([][]int, 100)
	ans = append(ans, []int{})
	ans = append(ans, []int{1})
	m[1] = 1
	ans[1] = []int{1}
	m[0] = 1
	for i := 2; i < 100; i++ {
		for j := 1; j < i; j++ {
			if m[i-j]*j > m[i] {
				m[i] = m[i-j] * j
				ans[i] = append(ans[i-j], j)
			}
			if (i-j)*j > m[i] {
				m[i] = (i - j) * j
				ans[i] = []int{i - j, j}
			}
		}
		fmt.Println("%d: %d\n", m[i], integerBreak(i))
	}
}
