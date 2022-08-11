package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	v := add(2, 3)
	fmt.Println(v)

	//	result, err := sqrt(16)
	result, err := sqrt(-16)
	if err != nil {
		fmt.Println((err))
	} else {
		fmt.Println(result)
	}

	//fmt.Println(sqrt(16))

}

// returns int different to regular programming lang
func add(x int, y int) int {
	z := x + y
	return z
}

// Go can return multiple values!!
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Invalid input no negative numbers")
	}
	return math.Sqrt(x), nil
}
