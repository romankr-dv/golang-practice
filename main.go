package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 5; i++ {
		fmt.Println("go"+"lang [", i, "]")
	}
	x := []float64{2.15, 3.14, 42.0, 29.5}
	fmt.Println(average(x))
	x1 := []float64{}
	fmt.Println(average(x1))

	s := "Hello, world"
	for i, r := range s {
		fmt.Printf("i%d r %c\n", i, r)
	}
	fmt.Println("----")
	a := []rune(s)
	for i, r := range a {
		fmt.Printf("i%d r %c\n", i, r)
	}
}

func average(x []float64) float64 {
	if len(x) == 0 {
		return 0
	}
	var total float64 = 0
	// for i := 0; i < len(x); i++ {
	for _, v := range x {
		total += v
	}
	return total / float64(len(x))
}
