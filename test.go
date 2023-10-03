package main

import (
	"fmt"
	"jcl/utilities"
)

func main() {
	z := map[rune]float64{
		'a': 0.5,
		'b': 0.30,
		'c': 0.19,
		'd': 0.01,
	}
	// hello.mydist
	x := utilities.Distribution{Distrib: z}
	fmt.Println(x.To_list())
}
