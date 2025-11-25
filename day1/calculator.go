package day1

import "fmt"

func Calculator() {
	var a, b float64
	var operator string
	fmt.Println("Enter the first number:")
	fmt.Scan(&a)
	fmt.Println("Enter the second number:")
	fmt.Scan(&b)
	fmt.Println("Enter operators (+,-,*,/):")
	fmt.Scan(&operator)

	switch operator {
	case "+":
		fmt.Printf("%f %s %f = %f", a, operator, b, a+b)

	case "-":
		fmt.Printf("%f %s %f = %f", a, operator, b, a-b)

	case "*":
		fmt.Printf("%f %s %f = %f", a, operator, b, a*b)

	case "/":
		if b == 0 {
			fmt.Printf("the number cant be divisible by 0")
		} else {
			fmt.Printf("%f %s %f = %f", a, operator, b, a/b)
		}

	}

}
