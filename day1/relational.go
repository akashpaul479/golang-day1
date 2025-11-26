package day1

import "fmt"

func Relation() {
	var a, b string
	fmt.Println("Enter the first value: ")
	fmt.Scan(&a)
	fmt.Println("Enter the second value: ")
	fmt.Scan(&b)
	fmt.Println(a < b)
	fmt.Println(a > b)
	fmt.Println(a <= b)
	fmt.Println(a >= b)
	fmt.Println(a == b)
	fmt.Println(a != b)
}
