package main

import (
	"fmt"
	"rsc.io/quote"
)

func main() {
	quote := quote.Go()
	fmt.Println("Hi ", quote)

	for i := 1; i < 5; i++ {
		fmt.Println("go"+"lang [", i, "]")
	}
}
