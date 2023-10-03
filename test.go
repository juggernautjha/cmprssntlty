package main

import (
	"fmt"
	"jcl/utilities"
)

func main() {
	z := map[string]float64{
		"A": 0.5,
		"B": 0.30,
		"C": 0.19,
		"D": 0.01,
	}
	// hello.mydist
	x := utilities.Distribution{Distrib: z}
	fmt.Println(x.To_list())
}
