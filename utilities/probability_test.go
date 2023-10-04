package utilities

import (
	"math"
	"testing"
)

func TestDistribution_Size(t *testing.T) {
	dist := Distribution{map[rune]float64{'a': 0.5, 'b': 0.25, 'c': 0.25}}
	size := dist.Size()
	if size != 3 {
		t.Errorf("Size() returned %d, expected %d", size, 3)
	}
}

func TestDistribution_Get_alphabet(t *testing.T) {
	dist := Distribution{map[rune]float64{'a': 0.5, 'b': 0.25, 'c': 0.25}}
	alphabet := dist.Get_alphabet()
	if len(alphabet) != 3 { //! Janky boi hehe
		t.Errorf("Get_alphabet() returned %d elements, expected %d", len(alphabet), 3)
	}
	if alphabet[0] != int32('a') || rune(alphabet[1]) != 'b' || alphabet[2] != int32('c') {
		t.Errorf("Get_alphabet() returned %v, expected ['a', 'b', 'c']", alphabet)
	}
}

func TestDistribution_To_list(t *testing.T) {
	dist := Distribution{map[rune]float64{'a': 0.5, 'b': 0.25, 'c': 0.25}}
	list := dist.To_list()
	if len(list) != 3 {
		t.Errorf("To_list() returned %d elements, expected %d", len(list), 3)
	}
	if math.Abs(list[0]-0.5) > 1e-9 || math.Abs(list[1]-0.25) > 1e-9 || math.Abs(list[2]-0.25) > 1e-9 {
		t.Errorf("To_list() returned %v, expected [0.5, 0.25, 0.25]", list)
	}
}

func TestDistribution_Entropy(t *testing.T) {
	dist := Distribution{map[rune]float64{'a': 0.5, 'b': 0.25, 'c': 0.25}}
	entropy := dist.Entropy()
	if math.Abs(entropy-1.5) > 1e-9 {
		t.Errorf("Entropy() returned %f, expected %f", entropy, 1.5)
	}
}
