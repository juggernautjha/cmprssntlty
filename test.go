package main

import (
	. "jcl/libraries"
)

func main() {
	// z := map[rune]float64{
	// 	'a': 0.5,
	// 	'b': 0.30,
	// 	'c': 0.19,
	// 	'd': 0.01,
	// }
	// // hello.mydist
	// x := utilities.Distribution{Distrib: z}
	// fmt.Println(x.To_list())

	x := BitMap{'a': "", 'b': "0", 'c': "10", 'd': "110"}
	h := Create_tree(x)
	j := Get_encoding(&h)
	print(j)

}
