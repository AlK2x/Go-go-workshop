package main

import (
	"fmt"
	"unicode/utf8"
	"strings"
	"sort"
)

func main() {
	if IsAnagramWithSort("анаГрамма", "ммааанагр") {
		fmt.Print("yes")
	} else {
		fmt.Print("no")
	}
}

func IsAnagramWithMap(first string, second string) bool {
	if utf8.RuneCount([]byte(first)) != utf8.RuneCount([]byte(second)) {
		return false
	}

	runesCount := utf8.RuneCount([]byte(first))
	mf := make(map[rune]int, runesCount)
	ms := make(map[rune]int, runesCount)
	for _, r := range strings.ToLower(first) {
		mf[r]++
	}

	for _, r := range strings.ToLower(second) {
		ms[r]++
	}

	for r, c := range mf {
		if c != ms[r] {
			return false;
		}
	}
	return true;
}

func IsAnagramWithSort(first string, second string) bool {
	fb := []byte(first)
	sb := []byte(second)
	if utf8.RuneCount(fb) != utf8.RuneCount(sb) {
		return false
	}

	a1 := []rune(strings.ToLower(first))
	a2 := []rune(strings.ToLower(second))
	sort.Slice(a1, func(i, j int) bool { return a1[i] < a1[j]})
	sort.Slice(a2, func(i, j int) bool { return a2[i] < a2[j]})

	for i := 0; i < len(a1); i++ {
		if a1[i] != a2[i] {
			return false
		}
	}

	return true
}