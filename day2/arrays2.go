package day2

import "fmt"

func Arrays2() {
	arr := []int{10, 20, 30, 40, 50}
	fmt.Println("Array is in reverse order:")
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Println(arr[i])
	}
}
