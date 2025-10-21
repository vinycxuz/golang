package fizzbuzz

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	testSuite := []struct {
		name   string
		input  int
		expect string
	}{
		{"ShouldReturnFizz",
			3,
			"Fizz",
		},
		{
			"ShouldReturnBuzz",
			5,
			"Buzz",
		},
		{
			"ShouldReturnFizzBuzz",
			15,
			"FizzBuzz",
		},
		{"ShouldReturnFizz",
			33,
			"Fizz",
		},
		{
			"ShouldReturnBuzz",
			55,
			"Buzz",
		},
		{
			"ShouldReturnFizzBuzz",
			150,
			"FizzBuzz",
		},
	}

	for _, tt := range testSuite {
		t.Run(tt.name, func(t *testing.T) {
			got := FizzBuzz(tt.input)

			assert.Equal(t, tt.expect, got)
		})
	}
}
