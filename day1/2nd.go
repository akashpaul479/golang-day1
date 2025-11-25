package day1

import "fmt"

func Exe1() {
	var a, b, c float64
	fmt.Println("Enter the three numbers:")
	fmt.Scan(&a, &b, &c)
	if a > b && a > c {
		fmt.Println("largest number is:", a)
	} else if b > a && b > c {
		fmt.Println("largest number is:", b)
	} else {
		fmt.Println("largest number is:", c)
	}

}
