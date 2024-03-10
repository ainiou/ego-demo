package utils

import (
	"testing"
)

func TestSliceToMap(t *testing.T) {
	a := []int{1, 2, 3}
	c := SliceToMap(a, func(b int) int {
		return b
	})
	t.Log(c)
}
