package main

import (
	"fmt"
)

func solve(s string, substr string) int {
	m := make(map[byte]int)
	for i := 0; i < len(substr); i++ {
		_, ok := m[substr[i]]
		if ok {
			m[substr[i]] = m[substr[i]] + 1
		} else {
			m[substr[i]] = 1
		}
	}
	
	//fmt.Println(m)
	start := 0
	end := -1
	for end < len(s)-1 {
		end += 1
		_, ok := m[s[end]]
		if ok {
			m[s[end]] = m[s[end]] - 1
		} else {
			m[s[end]] = -1
		}
		for m[s[end]] < 0 && start <= end {
			m[s[start]] += 1
			start += 1
		}
		if end - start + 1 == len(substr) {
			return start
		}
	}
	return -1
}

func main() {
	fmt.Println(solve("abcde e ", "bc"))  // 1
	fmt.Println(solve("abcde e ", "abc"))  // 0
	fmt.Println(solve("abcde e ", " bc"))  // -1
	fmt.Println(solve("Hello world", "w oorl")) // 3
}
