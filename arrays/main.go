package main

import "fmt"

func main() {
	var a [3]int
	fmt.Println(a) // [0, 0, 0]

	for i := range a {
		fmt.Println(i, a[i])
	}

	for i, v := range a {
		fmt.Println(i, v)
	}

	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}

	var b = [...]int{1, 2, 3, 4}
	fmt.Println(b) // [1, 2, 3, 4]

	var c = [...]string{"abc", "bcd"}
	fmt.Println(c) // ["abc", "bcd"]

	var d = [...]int{2: 3, 4: 3} // {0: 0, 1: 0, 2: 3, 3: 0, 4: 3}
	fmt.Println(d)

	var e = &a
	fmt.Println(e[0]) // 0

	var arr1 = [...]int{1, 2, 3, 4, 5}
	var arr2 = &arr1

	// Correct way to iterate over an array pointer
	for index, value := range *arr2 {
		arr2[index] += 1
		fmt.Println(index, value)
	}
}
