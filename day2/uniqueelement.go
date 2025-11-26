package day2

import "fmt"

func Uniqueele() {
	numbers := []int{1, 2, 2, 3, 4, 4, 4}
	slice := []int{}
	for _, m := range numbers {
		exists := false
		for _, n := range slice {
			if m == n {
				exists = true
				break
			}
		}
		if !exists {
			slice = append(slice, m)
		}
	}
	fmt.Println(slice)
}
