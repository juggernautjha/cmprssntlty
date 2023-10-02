package main

import (
	"fmt"
	"jcl/utilities"
)

// package main
//
// import "fmt"

// func main() {
// 	fmt.Println("Hello, world!")
// }

/*
Distribution will be endowed with the following functions:
validate -> verifies if it is indeed a valid probability distribution.
size -> pretty trivial, returns the size of the distribution. (In this case, the map)
get_alphabet -> returns the set of alphabets used in the distribution.
to_list -> converts the map to a list. go equivalent of [dist[s] for s in dist]

! God willing, I WILL write unit tests some day. I do not thing god wills that though.
*/

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
