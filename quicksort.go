package main

import (
	"time"
)

func Partition(low, high int, a []int) int {
	partitionPos := low - 1
	if low < high {
		pivot := a[high]
		// fmt.Println("\t low:", low)
		// fmt.Println("\t high:", high)
		// fmt.Println("\t pivot:", pivot)

		// search elements <= pivot
		for i := low; i < high; i++ {
			if a[i] <= pivot {
				// shift the partition position
				partitionPos++
				// swap val at i with val at partition position
				a[i], a[partitionPos] = a[partitionPos], a[i]
			}
		}
		// shift Partition position
		partitionPos++
		// swap pivot with val at partition position / high
		a[partitionPos], a[high] = a[high], a[partitionPos]

		// fmt.Println("\t partition:", a[partitionPos], "now set at position: ", partitionPos)
		// fmt.Println("\t after partition process: ", a)
	}
	return partitionPos
}

func Partition_OLD(low, high int, a []int) int {
	partitionPos := -1
	if low < high {
		pivot := a[low]
		i := low
		j := high

		// fmt.Println("\t low:", low)
		// fmt.Println("\t high:", high)
		// fmt.Println("\t pivot:", pivot)

		// Base rule
		for i < j {

			// Keep moving i towards right to find an element > pivot
			for a[i] <= pivot && i < high {
				// fmt.Println("\t i:", i, "shifting i to right -> ", i+1)
				i++
			}

			// Keep moving j towards left to find an element < pivot
			for a[j] > pivot && j > 0 {
				// fmt.Println("\t  j:", j, "shifting j to left: ", j-1)
				j--
			}

			// Swap both elements when found provided i < j
			if i < j {
				a[j], a[i] = a[i], a[j] // swap
				// fmt.Println("\t i < j so swapping i & j, now a:", a)
			}
		}

		// fmt.Println("\t current i:", i, " | j:", j)
		// fmt.Println("\t swapping pivot:", pivot, " with val at j:", a[j])
		// once base rule breaks, swap items at j and pivot
		a[j], a[low] = a[low], a[j] // swap
		// fmt.Println("\t partition:", a[j], "now set at position: ", j)
		partitionPos = j
	}
	// fmt.Println("\t after partition process: ", a)

	// return partition position
	return partitionPos
}

func QuickSort(low, high int, a []int) {
	// fmt.Println("\t --------------------")
	// fmt.Println("\t QuickSort -> ", a)
	if low < high {
		// fmt.Println("\t partitioning...")
		partitionPos := Partition(low, high, a)
		if partitionPos > -1 {
			QuickSort(low, partitionPos-1, a)
			QuickSort(partitionPos+1, high, a)
		}
	}
	// fmt.Println("\t --------------------")
}

func sortThisQuick(a []int) (time.Duration, []int) {
	start := time.Now()
	low := 0
	high := len(a) - 1
	// fmt.Println("input:", a)
	QuickSort(low, high, a)
	return time.Since(start), a
}
