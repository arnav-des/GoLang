package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	// Prompt the user to input a series of integers
	var n int
	fmt.Print("Enter the number of integers: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Enter integer #%d: ", i+1)
		fmt.Scan(&arr[i])
	}

	// Divide the array into 4 equal parts
	partSize := n / 4
	parts := make([][]int, 4)
	for i := 0; i < 4; i++ {
		start := i * partSize
		end := start + partSize
		if i == 3 {
			end = n
		}
		parts[i] = arr[start:end]
	}

	// Print the subarrays to be sorted by each goroutine
	for i := 0; i < 4; i++ {
		fmt.Printf("Sorting goroutine %d: %v\n", i+1, parts[i])
	}

	// Sort each subarray in a separate goroutine
	var wg sync.WaitGroup
	wg.Add(4)
	for i := 0; i < 4; i++ {
		go func(i int) {
			sort.Ints(parts[i])
			wg.Done()
		}(i)
	}
	wg.Wait()

	// Merge the sorted subarrays
	result := make([]int, n)
	copy(result, parts[0])
	for i := 1; i < 4; i++ {
		result = merge(result, parts[i])
	}

	// Print the sorted array
	fmt.Println("Sorted array:", result)
}

// Merge two sorted arrays into one sorted array
func merge(a, b []int) []int {
	result := make([]int, len(a)+len(b))
	i, j, k := 0, 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result[k] = a[i]
			i++
		} else {
			result[k] = b[j]
			j++
		}
		k++
	}
	for i < len(a) {
		result[k] = a[i]
		i++
		k++
	}
	for j < len(b) {
		result[k] = b[j]
		j++
		k++
	}
	return result
}
