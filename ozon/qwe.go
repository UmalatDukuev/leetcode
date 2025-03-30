package main

import (
	"fmt"
)

func main() {
	inc := func() {
		v++
	}

	v := 1
	fmt.Print(v, " ")
	inc()
	fmt.Print(v, " ")
}
