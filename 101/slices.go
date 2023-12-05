package main

import (
	"fmt"
	"slices"
)

func main() {
	var arr [5]int
	fmt.Println("first element", arr[0])

	var s []string
	fmt.Println("uninit", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "Mustafa"
	s[1] = "Kemal"
	s[2] = "Atatürk"
	fmt.Println("new Emp:", s, "len:", len(s), "cap:", cap(s))
	s = append(s, "Türkiye")
	fmt.Println("new Emp:", s, "len:", len(s), "cap:", cap(s))
	s = append(s, "Adana")
	fmt.Println("new Emp:", s, "len:", len(s), "cap:", cap(s))
	s = append(s, "Kilis")
	fmt.Println("new Emp:", s, "len:", len(s), "cap:", cap(s))
	s = append(s, "Ankara")
	fmt.Println("new Emp:", s, "len:", len(s), "cap:", cap(s))

	process_one := s[2:5] // 2 include , 5 exclude
	fmt.Println("s[2:5] : element 2 include, element 5 exclude:", process_one)

	process_two := s[:5]
	fmt.Println("s[:5] : element 5 and beyond exclude:", process_two)

	process_three := s[2:]
	fmt.Println("s[2:] : exclude until element 2 ", process_three)

	t1 := []string{"a", "l", "i"}
	t2 := []string{"a", "l", "i"}

	if slices.Equal(t1, t2) {
		fmt.Println("t1 and t2 are equal")
	}
}
