package main

import (
	"bytes"
	"fmt"
)

func main() {
	bt1 := []byte{'g', 'r', 'e', 'e', 'k', 's'}
	bt2 := []byte{'s', 'i', 'g', 'h', 't', 's'}

	fmt.Println("Original Runes")
	fmt.Printf("bt1 : %s \n", bt1)
	fmt.Printf("bt2 : %s \n", bt2)
	fmt.Println("To Bytes")
	fmt.Printf("bti : %v \n", bt1)
	fmt.Printf("bt2 : %v \n", bt2)
	fmt.Println("Converted Runes")
	fmt.Printf("bt1 : %v \n", bytes.ToUpper(bt1))
	fmt.Printf("bt2 : %v \n", bytes.ToUpper(bt2))
	fmt.Println("String version")
	fmt.Printf("bt1 : %s \n", bytes.ToUpper(bt1))
	fmt.Printf("bt2 : %s \n", bytes.ToUpper(bt2))
}
