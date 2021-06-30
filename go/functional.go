package main

import (
	"flag"
	"fmt"
	"math"
)

func main() {
	var powers int
	flag.IntVar(&powers, "power", 2, "raise a list to a power")
	flag.Parse()
	args := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	raiseToPowers := pow(powers)
	for _, value := range args {
		fmt.Printf("%d ", raiseToPowers(value))
	}
	fmt.Println()
}

// pow returns a function which raises a something to a power of p
func pow(p int) func(int) int {
	return func(base int) int {
		return int(math.Pow(float64(base), float64(p)))
	}
}
