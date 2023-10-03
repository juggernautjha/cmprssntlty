package utilities

import (
	"testing"
)

func TestDistribution_Size(t *testing.T) {
	dist := Get_distribution(map[rune]float64{'a': 0.5, 'b': 0.5})
	size := dist.Size()
	if size != 2 {
		t.Errorf("Expected size to be 2, but got %d", size)
	}
}

func TestDistribution_Get_alphabet(t *testing.T) {
	dist := Get_distribution(map[rune]float64{'a': 0.5, 'b': 0.5})
	alphabet := dist.Get_alphabet()
	if len(alphabet) != 2 {
		t.Errorf("Expected alphabet length to be 2, but got %d", len(alphabet))
	}
	if alphabet[0] != 'a' || alphabet[1] != 'b' {
		t.Errorf("Expected alphabet to be ['a', 'b'], but got %v", alphabet)
	}
}

func TestDistribution_To_list(t *testing.T) {
	dist := Get_distribution(map[rune]float64{'a': 0.5, 'b': 0.5})
	list := dist.To_list()
	if len(list) != 2 {
		t.Errorf("Expected list length to be 2, but got %d", len(list))
	}
	if list[0] != 0.5 || list[1] != 0.5 {
		t.Errorf("Expected list to be [0.5, 0.5], but got %v", list)
	}
}

func TestDistribution_Entropy(t *testing.T) {
	dist := Get_distribution(map[rune]float64{'a': 0.5, 'b': 0.5})
	entropy := dist.Entropy()
	if entropy != 1 {
		t.Errorf("Expected entropy to be 1, but got %f", entropy)
	}
}

func TestDistribution_Validate(t *testing.T) {
	// Valid distribution
	dist1 := Get_distribution(map[rune]float64{'a': 0.5, 'b': 0.5})
	if !dist1.validate() {
		t.Errorf("Expected distribution to be valid, but got invalid")
	}

	// Invalid distribution with values too small
	dist2 := Distribution{map[rune]float64{'a': 0.5, 'b': 0.0000001}}
	if dist2.validate() {
		t.Errorf("Expected distribution to be invalid, but got valid")
	}

	// Invalid distribution with sum not equal to 1
	dist3 := Distribution{map[rune]float64{'a': 0.5, 'b': 0.8}}
	if dist3.validate() {
		t.Errorf("Expected distribution to be invalid, but got valid")
	}
}
