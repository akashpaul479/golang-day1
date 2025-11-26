package day2

import "fmt"

func Arrays1() {
	var n [100]float64
	var total int
	fmt.Println("Enter no of elements:")
	fmt.Scan(&total)
	for i := 0; i < total; i++ {
		fmt.Println("enter the numbers:")
		fmt.Scan(&n[i])
	}
	for j := 1; j < total; j++ {
		if n[0] < n[j] {
			n[0] = n[j]
		}
	}
	fmt.Println("Enter the largest number:", n[0])
}
