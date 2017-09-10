package main

import (
	"fmt"
)

//Calc xx
func Calc(start, end int) (int, float32) {
	sumResult := 0
	needCount := make([]int, end-start)

	for index := start; index <= end; index++ {
		needCount = append(needCount, index)
	}

	for _, val := range needCount {
		sumResult += val
	}
	//var count float32
	//count = float32(end - start)
	return sumResult, float32(sumResult) / float32(end-start)
}

//Calc xx
func Calc2(vals ...int) (int, float32) {
	sumResult := 0

	for _, val := range vals {
		//if val%2＝=0
		//	continue
		sumResult += val
	}
	//var count float32
	//count = float32(end - start)
	return sumResult, float32(sumResult) / float32(len(vals))
}

func main() {
	fmt.Println(Calc(1, 100))
	fmt.Println(Calc2([]int{100, 200, 300}...))

}
