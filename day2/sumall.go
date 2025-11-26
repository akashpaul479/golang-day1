package day2

import "fmt"

func Allsum() {
	var n int
	fmt.Println("Enter the number:")
	fmt.Scan(&n)
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	fmt.Println("the sum off all number is:", sum)
}
