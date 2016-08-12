package main

import (
	"fmt"
	"strings"
)

func Index(vstr []string, t string) int {
	for i, v := range vstr {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(vstr []string, t string) bool {
	if Index(vstr, t) >= 0 {
		return true
	}
	return false
}

func Any(vstr []string, f func(string) bool) bool {
	for _, v := range vstr {
		if f(v) {
			return true
		}
	}
	return false
}

func All(vstr []string, f func(string) bool) bool {
	for _, v := range vstr {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vstr []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vstr {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Map(vstr []string, f func(string) string) []string {
	vsm := make([]string, len(vstr))
	for i, v := range vstr {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {
	strs := []string{"Foo", "bar", "Baz", "ZyX"}
	fmt.Println(Index(strs, "ZyX"))
	fmt.Println(Index(strs, "Zyx"))

	fmt.Println("Include ->", Include(strs, "Zyx"))

	fmt.Println("Any ->", Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "F")
	}))

	fmt.Println("All ->", All(strs, func(v string) bool {
		return strings.HasPrefix(v, "F")
	}))

	fmt.Println("Filter ->", Filter(strs, func(v string) bool {
		return strings.Contains(v, "a")
	}))

	fmt.Println("Map ->", Map(strs, func(v string) string {
		return strings.ToUpper(v)
	}))
}
