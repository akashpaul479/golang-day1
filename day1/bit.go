package day1

import "fmt"

func Bit() {
	var a, b int
	fmt.Println("Enter the first number")
	fmt.Scan(&a)
	fmt.Println("Enter the second number")
	fmt.Scan(&b)
	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	fmt.Println(a << 1)
	fmt.Println(a >> 1)
}
