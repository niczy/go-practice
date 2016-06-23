package main

import (
	"fmt"
)

func solve(s string, row int) string {
	ret := make([]string, row)
	i := 0
	d := 1
	for _, c := range s {
		nexti := i
		if i+d < row && i+d >= 0 {
			nexti = i + d
			fmt.Printf("nexti %d\n", nexti)
		} else {
			d = -d
			fmt.Printf("d %d\n", d)
			if i+d < row && i+d >= 0 {
				nexti = i + d
				fmt.Printf("nexti %d\n", nexti)
			}
		}
		fmt.Printf("i %d\n", i)
		ret[i] = ret[i] + string(c)
		i = nexti

	}
	rets := ""
	fmt.Printf("%v\n", ret)
	for i := 0; i < row; i++ {
		rets += ret[i]
	}
	return rets
}

func main() {
	fmt.Println(solve("PAYPALISHIRING", 1))
	fmt.Println(solve("PAYPALISHIRING", 3))
}
