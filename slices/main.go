package main

import "fmt"
func main() {
	var a []int // nil slice
	fmt.Println(a)

	var b = []int {}
	fmt.Println(b) // empty slice ( khac voi nil slice)

	// co 3 element in slice and len = 3 & cap = 3
	var c = []int {1, 2, 3}

	// co 2 elements in slice and len = 2 & cap = 3
	var d = c[:2]
	fmt.Println(d)

	// co 2 elements in slice and len = 2 & cap = 3
	var e = c[0:2:cap(c)]
	fmt.Println(e)

	// 2 elements in slice and len = 2, cap = 3
	var g = make ([]int, 2, 3)
	fmt.Println(g)

	// if cap of slice < 1024 then x2 cap
	// if cap of slice >= 1024 then x1.25 cap
	// when append(element)

	sl := make([]int, 1)
	printSlice(sl)

	for i := 0; i < 5; i++ {
		sl = myAppend(sl, i)
	}

	sl2 := []int{0}
	reverseAppend(sl2)
}

func printSlice(sl []int) {
	fmt.Printf("Slice %v\n", sl)
	fmt.Printf("len %v, Cap %v\n\n", len(sl), cap(sl))
}

func myAppend(sl []int, val int) []int {
	sl = append(sl, val)
	printSlice(sl)
	return sl
}

func reverseAppend(sl []int) {
	sl = append([]int{1, 2, 3}, sl...)
	fmt.Println(sl)
}