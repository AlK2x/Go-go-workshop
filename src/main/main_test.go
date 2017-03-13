package main

import "testing"

func TestIsAnagramWithSort(t *testing.T) {
	cases := []struct {
		first, second string
		out           bool
	} {
		{"Анаграмма", "граммаана", true},
		{"Hello,   ", "1", false},
		{"", "", true},
	}
	for _, c := range cases {
		got := IsAnagramWithSort(c.first, c.second)
		if got != c.out {
			t.Errorf("IsAnagram (%q, %q) : %v", c.first, c.second, c.out)
		}
	}
}

func TestIsAnagramWithMap(t *testing.T) {
	cases := []struct {
		first, second string
		out           bool
	} {
		{"Анаграмма", "граммаана", true},
		{"Hello,   ", "1", false},
		{"", "", true},
	}
	for _, c := range cases {
		got := IsAnagramWithMap(c.first, c.second)
		if got != c.out {
			t.Errorf("IsAnagram (%q, %q) : %v", c.first, c.second, c.out)
		}
	}
}

func BenchmarkIsAnagramWithSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsAnagramWithSort("Hello","Hello")
	}
}

func BenchmarkIsAnagramWithMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsAnagramWithMap("Hello","Hello")
	}
}