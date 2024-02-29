package main

import (
	"math/rand"
)

func getLargeSlice(count int) []int {
	a := []int{}

	for i := 0; i < count; i++ {
		b := rand.Intn(count)
		a = append(a, b)
	}
	return a
}
