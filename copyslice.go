//Copy one slice into another
package main

import "fmt"

func main() {

	slice6 := []string{"Greek", "For", "Greek", "GFG"}
	slice7 := make([]string, 3)
	fmt.Println(slice6)
	fmt.Println(slice7)

	copy(slice7, slice6)
	fmt.Println(slice7)
	// slice1 := []int{58, 69, 40, 45, 11, 56, 67, 21, 65}
	// var slice2 []int
	// slice3 := make([]int, 5)
	// slice4 := []int{78, 50, 67, 77}

	// fmt.Println(slice1)
	// fmt.Println(slice2)
	// fmt.Println(slice3)
	// fmt.Println(slice4)

	// fmt.Println("Start the copy")

	// copy1 := copy(slice2, slice1)
	// copy2 := copy(slice3, slice1)
	// copy3 := copy(slice4, slice1)

	// fmt.Println(copy1)
	// fmt.Println(copy2)
	// fmt.Println(copy3)

	// fmt.Println(slice2)
	// fmt.Println(slice3)
	// fmt.Println(slice4)
	// // fmt.Println(slice2)
	// // fmt.Println(slice3)
	// // fmt.Println(slice4)

}
