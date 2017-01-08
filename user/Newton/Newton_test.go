package main

import (
	"testing"
	"math"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{Sqrt(2), math.Sqrt(2)},
		{Sqrt(3), math.Sqrt(3)},
	}
	for _, c := range cases {
		if c.in - c.want > 0.0000001 {
			t.Errorf("Error")
		}
	}
}