package random

import (
	"testing"

)

func TestNewRandomString(t *testing.T) {
	cases := []int{10, 100, 1000}
	for _, c := range cases {
		randStr := NewRandomString(c)

		if len(randStr) != c {
			t.Errorf("wrong length: expected %d, got %d", c, len(randStr))
		}
	}
}