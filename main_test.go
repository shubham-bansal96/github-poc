package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		val1     int
		val2     int
		expected int
	}{
		{
			name:     "add 4 & 5",
			val1:     4,
			val2:     5,
			expected: 9,
		},
		{
			name:     "add -2 & 3",
			val1:     -2,
			val2:     3,
			expected: 1,
		},
		{
			name:     "add 3 & 0",
			val1:     3,
			val2:     0,
			expected: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(testing *testing.T) {
			res := add(test.val1, test.val2)
			assert.EqualValues(t, res, test.expected)
		})
	}
}
