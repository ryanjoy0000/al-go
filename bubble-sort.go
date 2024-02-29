package main

import (
	"time"
)

func bubbleSort(arr []int) ([]int, time.Duration) {
	start := time.Now()
	// fmt.Println("input: ", arr)
	for j := len(arr) - 1; j >= 0; j-- {
		// fmt.Println("j:", j)
		for i := 0; i < j; i++ {
			// fmt.Println("\ti:", i)
			if arr[i] > arr[i+1] {
				// fmt.Println("\tswapping ", arr[i], " and ", arr[i+1])
				temp := arr[i+1]
				arr[i+1] = arr[i]
				arr[i] = temp
			}
		}
	}
	return arr, time.Since(start)
}
