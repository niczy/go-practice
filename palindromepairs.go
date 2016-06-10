package main

import (
	"fmt"
)

func isPalindrom(bytes []byte) bool {
	for i := 0; i < len(bytes)/2; i++ {
		if bytes[i] != bytes[len(bytes)-i-1] {
			return false
		}
	}
	return true
}

func revers(s string) string {
	bytes := []byte(s)
	for i := 0; i < len(s)/2; i++ {
		bytes[i], bytes[len(bytes)-i-1] = bytes[len(bytes)-i-1], bytes[i]
	}
	return string(bytes)
}

func findPairs(words []string) [][]int {
	ret := make([][]int, 0)
	m := make(map[string]int)
	for i, s := range words {
		m[revers(s)] = i
	}
	for idx, s := range words {
		bytes := []byte(s)
		for i := 0; i < len(bytes); i++ {
			ss := string(bytes[i:])
			jdx, ok := m[ss]
			if isPalindrom(bytes[0:i]) && ok && jdx != idx {
				ret = append(ret, []int{idx, jdx})
			}
		}
	}
	return ret
}

func main() {
	fmt.Println(findPairs([]string{"bat", "tab", "cat"}))
	fmt.Println(findPairs([]string{"abcd", "dcba", "lls", "s", "sssll"}))
}
