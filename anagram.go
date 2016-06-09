package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func isAnagram(s string, t string) bool {
	a := strings.Split(s, "")
	b := strings.Split(t, "")
	sort.Strings(a)
	sort.Strings(b)
	return reflect.DeepEqual(a, b)
}

func main() {
	fmt.Println(isAnagram("ab", "ba"))
}
