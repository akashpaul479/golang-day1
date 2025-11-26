package day2

import "fmt"

func Evenloop() {
	var n int
	fmt.Println("Enter the number")
	fmt.Scan(&n)
	for i := 0; i <= n; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
