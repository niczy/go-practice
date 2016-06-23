package main

import (
	"fmt"
)

const (
	BUY      = 0
	SELL     = 1
	COOLDOWN = 2
)

func solve(prices []int) {
	fmt.Printf("prices are: %v\n", prices)
	n := len(prices)
	f := make([][3]int, n)
	p := make([][3]int, n)
	q := make([][3]int, n)
	for i := 1; i < len(prices); i++ {
		f[i] = [3]int{-1, -1, -1}
		for j := i - 1; j >= 0; j-- {
			if f[j][COOLDOWN] > f[i][BUY] {
				f[i][BUY] = f[j][COOLDOWN]
				p[i][BUY] = j
				q[i][BUY] = COOLDOWN
			}
			if f[j][BUY]+prices[i]-prices[j] > f[i][SELL] {
				f[i][SELL] = f[j][BUY] + prices[i] - prices[j]
				p[i][SELL] = j
				q[i][SELL] = BUY
			}
			if f[j][SELL] > f[i][COOLDOWN] {
				f[i][COOLDOWN] = f[j][SELL]
				p[i][COOLDOWN] = j
				q[i][COOLDOWN] = SELL
			}
			if f[j][COOLDOWN] > f[i][COOLDOWN] {
				f[i][COOLDOWN] = f[j][COOLDOWN]
				p[i][COOLDOWN] = j
				q[i][COOLDOWN] = COOLDOWN
			}

		}
	}
	d := SELL
	if f[n-1][COOLDOWN] > f[n-1][SELL] {
		d = COOLDOWN
	}
	res := make([]string, 0)
	i := n - 1
	for {
		fmt.Printf("i %d, d %d\n", i, d)
		switch d {
		case BUY:
			res = append(res, "buy")
			break
		case SELL:
			res = append(res, "sell")
			break
		case COOLDOWN:
			res = append(res, "cooldown")
			break
		}
		if i == 0 {
			break
		}

		nexti := p[i][d]
		if i-nexti > 1 {
			for j := 0; j < i-nexti-1; j++ {
				res = append(res, "cooldown")
			}
		}
		d = q[i][d]
		i = nexti
	}
	n = len(res)
	fmt.Printf("f: %v\n", f)
	for k := 0; k < len(res)/2; k++ {
		res[k], res[n-k-1] = res[n-k-1], res[k]
	}
	fmt.Printf("res: %v\n", res)
}

func main() {
	solve([]int{1, 2, 3, 0, 2})
	solve([]int{4, 1, 3, 9, 0, 2})
	solve([]int{4, 1, 3, 9, 0, 2, 100})
}
