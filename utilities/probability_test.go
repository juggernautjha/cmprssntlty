// BEGIN: abpxx6d04wxr
package utilities

import (
	"math"
	"reflect"
	"testing"
)

func TestDistribution_Validate(t *testing.T) {
	testCases := []struct {
		name     string
		input    Distribution
		expected bool
	}{
		{
			name: "Valid distribution",
			input: Distribution{
				Distrib: map[string]float64{
					"a": 0.5,
					"b": 0.25,
					"c": 0.25,
				},
			},
			expected: true,
		},
		{
			name: "Invalid distribution",
			input: Distribution{
				Distrib: map[string]float64{
					"a": 0.5,
					"b": 0.25,
					"c": 0.26,
				},
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.input.validate()
			if actual != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, actual)
			}
		})
	}
}

func TestGet_fromfrequency(t *testing.T) {
	input := map[string]int64{"a": 1, "b": 2, "c": 3}
	expectedOutput := Distribution{map[string]float64{"a": 0.16666666666666666, "b": 0.3333333333333333, "c": 0.5}}

	output := Get_fromfrequency(input)

	if !reflect.DeepEqual(output, expectedOutput) {
		t.Errorf("Test failed: expected %v, but got %v", expectedOutput, output)
	}
}

func TestDistribution_Size(t *testing.T) {
	testCases := []struct {
		name     string
		input    Distribution
		expected int
	}{
		{
			name: "Valid distribution",
			input: Distribution{
				Distrib: map[string]float64{
					"a": 0.5,
					"b": 0.25,
					"c": 0.25,
				},
			},
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.input.Size()
			if actual != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, actual)
			}
		})
	}
}

func TestDistribution_Get_alphabet(t *testing.T) {
	tests := []struct {
		name   string
		dist   Distribution
		result []string
	}{
		{
			name: "Empty distribution",
			dist: Distribution{
				Distrib: map[string]float64{},
			},
			result: []string{},
		},
		{
			name: "Non-empty distribution",
			dist: Distribution{
				Distrib: map[string]float64{
					"a": 0.5,
					"b": 0.25,
					"c": 0.25,
				},
			},
			result: []string{"a", "b", "c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dist.Get_alphabet(); !equal(got, tt.result) {
				t.Errorf("Distribution.Get_alphabet() = %v, want %v", got, tt.result)
			}
		})
	}
}

func TestDistribution_To_list(t *testing.T) {
	tests := []struct {
		name   string
		dist   Distribution
		result []float64
	}{
		{
			name: "Empty distribution",
			dist: Distribution{
				Distrib: map[string]float64{},
			},
			result: []float64{},
		},
		{
			name: "Non-empty distribution",
			dist: Distribution{
				Distrib: map[string]float64{
					"a": 0.5,
					"b": 0.25,
					"c": 0.25,
				},
			},
			result: []float64{0.5, 0.25, 0.25},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dist.To_list(); !equalFloat(got, tt.result) {
				t.Errorf("Distribution.To_list() = %v, want %v", got, tt.result)
			}
		})
	}
}

func TestDistribution_Entropy(t *testing.T) {
	tests := []struct {
		name   string
		dist   Distribution
		result float64
	}{
		{
			name: "Empty distribution",
			dist: Distribution{
				Distrib: map[string]float64{},
			},
			result: 0,
		},
		{
			name: "Non-empty distribution",
			dist: Distribution{
				Distrib: map[string]float64{
					"a": 0.5,
					"b": 0.25,
					"c": 0.25,
				},
			},
			result: 1.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dist.Entropy(); math.Abs(got-tt.result) > 1e-7 {
				t.Errorf("Distribution.Entropy() = %v, want %v", got, tt.result)
			}
		})
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func equalFloat(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if math.Abs(v-b[i]) > 1e-7 {
			return false
		}
	}
	return true
}
