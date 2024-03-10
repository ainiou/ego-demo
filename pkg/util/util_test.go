package utils

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	result := Slice(slice, 3)
	fmt.Println(result)
}

func TestDurationToZero(t *testing.T) {
	duration := DurationToZero()
	fmt.Println(duration.Seconds())
}
