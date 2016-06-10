package main

import (
	"fmt"
)

func main() {
	arr := []int{0, 1, 2, 0, 3, 4, 0, 4}
	fmt.Println(arr)
	i := 0
	j := 0
	n := len(arr)
	for j < n {
		if arr[j] != 0 {
			arr[i] = arr[j]
			i += 1
		}
		j += 1
	}
	for ; i < n; i++ {
		arr[i] = 0
	}
	fmt.Println(arr)
}
