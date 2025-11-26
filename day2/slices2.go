package day2

import "fmt"

func Slices3() {
	names := []string{"akash", "abhi", "muzz", "vishal"}
	remove := 3
	newnames := append(names[:remove], names[remove+1:]...)
	fmt.Println(newnames)
}
