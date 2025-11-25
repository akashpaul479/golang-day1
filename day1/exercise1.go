package day1

import "fmt"

func Exe() {
	var name string
	var age int
	fmt.Println("Enter the name:")
	fmt.Scan(&name)
	fmt.Println("Enter your age:")
	fmt.Scan(&age)
	fmt.Println("Hello", name, "you are", age, "years old")

}
