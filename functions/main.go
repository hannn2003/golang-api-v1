package main

import "fmt"

// function in golang hỗ trợ số lượng tham số thay đổi, biến số lượng tham số phải là tham số cuối cùng và biến này phải là kiểu slice

// name function
func Add_fn(a, b int) int {
	return a + b
}

func Swap(a, b int) (int, int) {
	return b, a
}

func Sum(a int, more ...int) int {
	for _, v:= range more {
		a += v
	}
	return a
}

func Print(a ...interface{}) {
	fmt.Println(a...)
}

// truyền vào array sẽ giúp
// nội dung của biến x không bị thay đổi
func once(x [3]int) {
	for i := range x {
		x[i] *= 2
	}
}

func twice(x []int) {
	for i := range x {
		x[i] *= 2
	}
}

func main() {

// anonymous function
	var Add = func(a, b int) int {
		return a + b
	}

	fmt.Println(Add(1, 2))

	fmt.Println(Swap(1, 2))

	sl := []int {1, 2, 3, 4}
	fmt.Println(Sum(1, sl...))

	var a = []interface{}{123, "abc"}

	Print(a) // [123, abc]

	Print(a...) // 123 abc

	data := [...]int{1, 2, 3}

	once(data)

	fmt.Println(data) // [1 2 3]

	twice(data[0:])
	fmt.Println(data) // [2 4 6]
}