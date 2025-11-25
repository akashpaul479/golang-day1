package day1

import "fmt"

func Statements() {
	var name string
	fmt.Print("Enter the name:")
	fmt.Scan(&name)
	if len(name) <= 10 {
		fmt.Println("name", name)
	} else {
		fmt.Println("maximum letter exceeds")
	}
}
