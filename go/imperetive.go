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
	for _, value := range args {
		fmt.Printf("%d ", int(math.Pow(float64(value), float64(powers))))
	}
	fmt.Println()
}
