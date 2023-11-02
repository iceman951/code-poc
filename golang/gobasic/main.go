package main

import (
	"basic/customer"
	"fmt"
	"unicode/utf8"
)

// must use main function only
// ; is optional

// package or global scope declaration
// var y = 10

func main() {
	println("Hello World2")

	// --- fmt
	fmt.Printf("Hello %T\n", true)
	fmt.Printf("Hello %#v\n", 10)

	// --- short declaration
	x := 10
	println(x)

	// --- if
	score := 50
	if score >= 50 || score <= 100 {
		println("Pass")
	} else if score >= 20 {
		println("Fail")
	}

	// array
	//var z [3]int = [3]int{1, 2}
	// short declaration
	// z := [...]int{1, 2, 3, 4} array size = input into array
	// z := []int{1, 2, 3} <= slice , can be add member with append
	// append(z, 4) ** ไม่ได้แก้ไขค่า z แต่ apppend ไปเลข 4 ใน array แล้ว return ค่ากลับมา
	z := [3]int{1, 2}
	fmt.Println(z)

	name := "ไอซ์"
	println(utf8.RuneCountInString(name))

	//map *java c# call Doctionary
	countries := map[string]string{}
	countries["th"] = "Thailand"
	countries["us"] = "United State"

	country, ok := countries["jp"]

	println(countries["th"])

	if !ok {
		println("no value")
	} else {
		println(country)
	}

	values := []int{10, 20, 30}

	for i := 0; i < len(values); i++ {
		println(values[i])
	}

	for _, v := range values {
		println(v)
	}

	sum, text := sum(99, 1)
	fmt.Printf("sum: %v %v\n", sum, text)


	println(customer.Name)
}

func sum(a int, b int) (int, string) {
	return a + b, "Hello"
}
