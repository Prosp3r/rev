package main

import (
	"fmt"
	"math/rand"
	"time"
)

/* Declaration
c := make(chan int)
var c chan int
c = make(chan int)
*/

func main() {
	//c := make(chan string)
	// go boring("borin!!!", c)
	// c := boring("Boring!")
	c := fanIn(boring("Boring A!!!"), boring("Boring B!!!"), boring("Boring C!!!"), boring("Borind D!!!"))

	for i := 0; i < 100; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}

	// time.Sleep(time.Duration(1000) * time.Millisecond)
	fmt.Println("You're boring, I'm leaving")
}

//Generator pattern - returns a channel
// func boring(msg string) <-chan string { //Generator
// 	c := make(chan string)
// 	go func() {
// 		for i := 0; i < 100; i++ {
// 			c <- fmt.Sprintf("%s %d", msg, i)
// 			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
// 		}
// 	}()
// 	return c
// }
func boring(input string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 100; i++ {
			c <- fmt.Sprintf("%s %d", input, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

// func fanIn(input1, input2 <-chan string) <-chan string {
// 	c := make(chan string)

// 	go func() {
// 		for {
// 			c <- <-input1
// 		}
// 	}()

// 	go func() {
// 		for {
// 			c <- <-input2
// 		}
// 	}()
// 	return c
// }
//variadic fanIn function
func fanIn(input1 ...<-chan string) <-chan string {
	c := make(chan string)

	for f := 0; f < len(input1)-1; f++ {
		go func() {
			for {
				c <- <-input1[f]
			}
		}()
	}
	return c
}
