package main

import "fmt"

func main() {
	arr := []int32{12, 11, 13, 5, 6}
	fmt.Println(arr)
	for a := 1; a < len(arr); a++ {
		j := a

		for j > 0 {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
			j = j - 1
		}
	}
	fmt.Println(arr)
}
