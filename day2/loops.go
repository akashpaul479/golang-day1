package day2

import (
	"fmt"
)

func Loop1() {
	var n int
	fmt.Println("Enter the number")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Println(i)
	}
}
