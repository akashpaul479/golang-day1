package day2

import "fmt"

func Loops5() {
	var n int
	fmt.Println("Enter the number:")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
