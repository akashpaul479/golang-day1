package day2

import "fmt"

func Slices1() {
	numbers := []int{20, 40, 60, 80, 100}
	for _, numbers := range numbers {
		if numbers >= 50 {
			fmt.Println(numbers)
		}
	}

}
