package main

import (
	"fmt"
)

func reverseVowels(s string) string {
	bytes := []byte(s)
	i := 0
	j := len(s) - 1
	m := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	for i < j {
		for {
			_, ok := m[bytes[i]]
			if !ok {
				i += 1
			} else {
				break
			}
			if i >= j {
				break
			}
		}
		for {
			_, ok := m[bytes[j]]
			if !ok {
				j -= 1
			} else {
				break
			}
			if j <= i {
				break
			}
		}
		if i != j {
			bytes[i], bytes[j] = bytes[j], bytes[i]
		}
		fmt.Println(i, j)
		i += 1
		j -= 1

	}
	return string(bytes)
}
func main() {
	fmt.Println(reverseVowels("hello"))
	fmt.Println(reverseVowels("leetcode"))
}
