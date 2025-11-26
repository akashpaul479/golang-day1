package day2

import "fmt"

func Oddloop() {
	var n int
	fmt.Println("enter the number")
	fmt.Scan(&n)
	for i := 0; i <= 50; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}
