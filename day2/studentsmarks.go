package day2

import "fmt"

func Smarks() {
	var n int
	fmt.Println("Enter number of students:")
	fmt.Scan(&n)
	marks := make([]int, n)

	fmt.Println("Enter the marks:")
	for i := 0; i < n; i++ {
		fmt.Scan(&marks[i])
	}
	fmt.Println("All marks:")
	for _, m := range marks {
		fmt.Println(m)

	}
	total := 0
	highest := marks[0]
	lowest := marks[0]

	for _, m := range marks {
		total += m
		if m > highest {
			highest = m
		}
		if m < lowest {
			lowest = m
		}
	}
	average := total / n

	fmt.Println("total value is:", total)
	fmt.Println("Highest value is:", highest)
	fmt.Println("Lowest value is:", lowest)
	fmt.Println("Average value is:", average)

}
