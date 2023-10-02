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

type distributionMap[K comparable] map[K]float64

type Distribution struct {
	Distrib distributionMap[string] // DO not imagine using something else. I believe I am fucking up, but who knows.
}

/*
Distribution will be endowed with the following functions:
validate -> verifies if it is indeed a valid probability distribution.
size -> pretty trivial, returns the size of the distribution. (In this case, the map)
get_alphabet -> returns the set of alphabets used in the distribution.
to_list -> converts the map to a list. go equivalent of [dist[s] for s in dist]

! God willing, I WILL write unit tests some day. I do not thing god wills that though.
*/

func Get_distribution(t distributionMap[string]) Distribution {
	// Always use this method to initialize a new distribution. This also checks for validity.
	temp_dist := Distribution{t}
	if temp_dist.validate() {
		return temp_dist
	}
	log.Fatal("Invalid distribution")
	return temp_dist
}

func (dist *Distribution) validate() bool {
	//Stupid rip-off of the SCl validate method. just checks if the the distribution is valid.
	var sum float64 = 0
	for _, v := range dist.Distrib {
		if v <= 1e-6 {
			log.Fatal("Invalid probability distribution")
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

func (dist *Distribution) Get_alphabet() []string {
	//fun fact, we don't actually need a set because maps do not allow duplicate keys.
	var temp []string
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
