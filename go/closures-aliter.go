package main

import "fmt"

func main() {
	s := make([]func(), 4)
	for i := 0; i < 4; i++ {
		// closure-capture/pinning i locally to the loop construct
		// now the output won't be `4, 4, 4, 4` as each of them have a different i
		i := i
		s[i] = func() {
			fmt.Printf("%d @ %p\n", i, &i)
		}
	}
	// fmt.Println(i)
	// I cannot do the above, but the value of i somewhere is 4
	for i := 0; i < 4; i++ {
		s[i]()
	}
}
