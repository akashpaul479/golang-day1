package day1

import "fmt"

func Logical() {
	var a, b, c int
	fmt.Println("Enter the first number:")
	fmt.Scan(&a)
	fmt.Println("Enter the first number:")
	fmt.Scan(&b)
	fmt.Println("Enter the first number:")
	fmt.Scan(&c)
	fmt.Println((a < b) && (a > c))
	fmt.Println((a < b) || (a > c))
	fmt.Println(!(a < b))
}
