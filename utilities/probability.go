/*
For doing `things` with probability distributions.
Author: Jha
Yeah, screw me, I gave in. Having a class for probability distributins was too good to not do.
*/
package utilities

import (
	"log"
	"math"
)

/*
Distribution will be endowed with the following functions:
validate -> verifies if it is indeed a valid probability distribution.
size -> pretty trivial, returns the size of the distribution. (In this case, the map)
get_alphabet -> returns the set of alphabets used in the distribution.
to_list -> converts the map to a list. go equivalent of [dist[s] for s in dist]

! God willing, I WILL write unit tests some day. I do not thing god wills that though.
*/
type Distribution struct {
	Distrib map[rune]float64 // DO not imagine using something else. I believe I am fucking up, but who knows.
}

func Get_distribution(t map[rune]float64) Distribution {
	// Always use this method to initialize a new distribution. This also checks for validity.
	temp_dist := Distribution{t}
	if temp_dist.validate() {
		return temp_dist
	}
	log.Fatal("Invalid distribution")
	return temp_dist
}

func Get_fromfrequency(t map[rune]int64) Distribution {
	var sum float64 = 0
	for _, v := range t {
		sum += float64(v)
	}

	temp_map := make(map[rune]float64)
	for i, j := range t {
		temp_map[i] = float64(j) / sum
	}
	return Distribution{temp_map}
}

func (dist *Distribution) validate() bool {
	//Stupid rip-off of the SCl validate method. just checks if the the distribution is valid.
	var sum float64 = 0
	for _, v := range dist.Distrib {
		if v <= 1e-6 {
			return false
		}
		sum += v
	}
	if math.Abs(sum-1) <= 1e-7 {
		return true
	} else {
		return false
	}
}

func (dist *Distribution) Size() int {
	//returns the size of the distribution. can i get more boring?
	return len(dist.Distrib)
}

func (dist *Distribution) Get_alphabet() []rune {
	//fun fact, we don't actually need a set because maps do not allow duplicate keys.
	var temp []rune
	for k := range dist.Distrib {
		temp = append(temp, k)
	}
	return temp
}

func (dist *Distribution) To_list() []float64 {
	//returns the distribution as a list. pretty trivial.
	var temp []float64
	for _, v := range dist.Distrib {
		temp = append(temp, v)
	}
	return temp
}

/*
Now for the more intense and memory using functions:
Entropy -> places a lower bound on the number of bits needed on average to encode the symbols of a message.
*/
func (dist *Distribution) Entropy() float64 {
	var entropy float64 = 0
	for _, val := range dist.Distrib {
		if val == 0 {
			continue
		}
		entropy += -val * math.Log2(val)
	}
	return entropy
}

/*
Notes: I did not implement the CDF because
1. It is fucking trivial, and
2. I do not think I will use it.
*/
