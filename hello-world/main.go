package main

import (
	"fmt"
	"maps"
	"slices"
)

// function
func plus(x, y, z int) int{
	return x + y + z
}

// function: multiple return values
func multipleReturnValues() (int, string) {
	return 1, "2"
}

// variadic function 
func sum(nums ...int) int{
	total := 0
	for _, num := range nums {
		total += num
	}

	return total
}

// closures
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// pointers
func zeroVal(ival int) {
	ival = 0
} 

func zeroptr(iptr *int) {
	*iptr = 0
}

// struct
type Product struct {
	name string
	price float32
}

func updatePrice(p *Product, newPrice float32){
	p.price = newPrice
}
func main() {
	fmt.Println("Hello World")

	res := plus(1, 2, 3)
	fmt.Println("res", res)

	vir1, vir2 := multipleReturnValues()
	fmt.Println("multipleReturnValues", vir1, vir2)

	nums := []int{1, 2, 3, 4}
	fmt.Println("variadic function response:", sum(nums...))

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
  fmt.Println(newInts())

	// Variables
	var v = 1;
	v = 2;
	fmt.Println(v)

	// Constants
	// const c = 1; c can't be different from 1
	

	// for loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		// if - else
		if i % 2 == 0 {
			fmt.Println("i is divisible by 2", i)
		}
		i = i + 1
	}

	// Array
	var a [5]int
	a[4] = 100
	fmt.Println(a) // default [0, 0, 0, 0, 100]

	var people = []string{"A", "B", "C", "D"}
	fmt.Println("people", people)
	for _, value := range people {
		switch  value {
			case "A":
				fmt.Println("His name is A")
			case "B":
				fmt.Println("Her name is B")
			case "C":
				fmt.Println("His name is C")
			default:
				fmt.Println("Can't find name")
		}
	}

	// Slice
	var s []string // s == nil and len(s) == 0 is true
	fmt.Println("unInit:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("unInit:", s, s == nil, len(s) == 0)
	fmt.Println("emp: ", s, "len: ", len(s), cap(s)) // len is length and cap is capacity of s
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)	

	s2 := make([]string, len(s))
	fmt.Println("len(s2): ", len(s2)) // length = 6
	copy(s2, s)
	fmt.Println("s2", s2) // s2 as s

	s3 := s[0:]
	fmt.Println("s3", s3)
	if slices.Equal(s2, s3) {
		fmt.Println("s2 == s3")
	}else {
		fmt.Println("s2 != s3")
	}

	// Maps
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 10
	fmt.Println("map m: ", m)
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map n:", n)
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
	
	// range over built-in Types
	numbers := []int{1,2,3}
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Println("sum: ", sum)

	//pointer
	p := 1
	zeroVal(p)
	fmt.Println("p", p)
	zeroptr(&p)
	fmt.Println("p", p)

	product := Product{name: "Product 1", price: 100}
	updatePrice(&product, 120)
	fmt.Println("new price", product.price)
}