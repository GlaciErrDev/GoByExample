package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("empty string :", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append:", s)

	c := make([]string, len(s))
	fmt.Println("new empty:", c)
	copy(c, s)
	fmt.Println("copied:", c)

	sl1 := s[2:5]
	fmt.Println("sl1:", sl1)

	sl2 := s[:5]
	fmt.Println("sl2:", sl2)

	sl3 := s[2:]
	fmt.Println("sl3:", sl3)

	t := []string{"g", "h", "i"}
	fmt.Println("declare string array:", t)

	twoD := make([][]int, 3)
	for i := 0; i < len(twoD); i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d:", twoD)

}
