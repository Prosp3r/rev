package main

import (
	"fmt"
	"time"

	"github.com/rev/utility"
)

func main() {
	// arr := []int32{5, 1, 4, 2, 8}
	arr := utility.ArrayGen(10)

	//sort
	count := int32(0)
	fmt.Println(arr)
	startTime := time.Now().Nanosecond()
	fmt.Println(startTime)
	arrs, count := bsort(arr, count)
	fmt.Printf("Final array :\n%v \n", arrs)
	endTime := time.Now().Nanosecond()
	fmt.Println(endTime)
	fmt.Println((endTime - startTime) / 1000000)
}

func bsort(arr []int32, count int32) ([]int32, int32) {
	// count := 0
	for x := 0; x < len(arr)-1; x++ {
		if arr[x] > arr[x+1] {
			arr[x], arr[x+1] = arr[x+1], arr[x]
			count++
		}
	}
	fmt.Println(arr)

	if count > 0 {
		bsort(arr, 0)
	}
	return arr, count
}
