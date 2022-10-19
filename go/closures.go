package main

import "fmt"

func perform(d func()) {
	d()
}

func main() {
	for i := 0; i < 4; i++ {
		// i := i
		v := func() {
			fmt.Printf("%d @ %p\n", i, &i)
		}
		// (Q) why is the address of i same, if v is a different closure, shouldn't I get a new closure of v with a new instance of i?
		// (A) v is a closure over &i, if we pin this, it'll be different, without the pinning there'd be ONE closure with ONE environment which will be references to the value `i` and at the time the closure is called, the value will be i = 4
		perform(v)
	}
}
