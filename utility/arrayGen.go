package utility

import (
	"math/rand"
	"time"
)

func ArrayGen(lenn int32) []int32 {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int32, lenn)
	for x := int32(0); x < lenn; x++ {
		arr[x] = rand.Int31n(100) //rand(1, 1800)
	}
	return arr
}
