package day1

import "fmt"

func EvenOdd() {
	var number int
	fmt.Println("Enter the number:")
	fmt.Scan(&number)
	if number%2 == 0 {
		fmt.Println("it is even")
		fmt.Println("and the number is:", number)
	} else {
		fmt.Println("it is odd")
		fmt.Println("and the number is:", number)
	}

}
