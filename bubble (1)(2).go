/*
Write a Bubble Sort program in Go. The program should prompt the user to type in a sequence of up to 10 integers.
The program should print the integers out on one line, in sorted order, from least to greatest.

Use your favorite search tool to find a description of how the bubble sort algorithm works.
As part of this program, you should write a function called BubbleSort() which takes a slice of integers as an argument and returns nothing.
The BubbleSort() function should modify the slice so that the elements are in sorted order.
A recurring operation in the bubble sort algorithm is the Swap operation which swaps the position of two adjacent elements in the slice.

You should write a Swap() function which performs this operation. Your Swap() function should take two arguments,
a slice of integers and an index value i which indicates a position in the slice. The Swap() function should return nothing,
but it should swap the contents of the slice in position i with the contents in position i+1
*/

package main

import "fmt"

func main() {
	var n int = 10

	fmt.Println("Enter an integer please:")

	sli := make([]int, n)

	poblateSlice(sli, 0, n)
	bubbleSort(sli)
}

func poblateSlice(sli []int, ip, n int) {

	if n == 0 {
		return
	}

	if num, err := getNumber(&sli[ip]); num != 1 {
		panic(err)
	}

	poblateSlice(sli, ip+1, n-1)
}

func getNumber(num *int) (int, error) {
	return fmt.Scan(num)
}

func bubbleSort(numbers []int) {
	for i := len(numbers); i > 0; i-- {
		swap(numbers, i)
	}

	fmt.Println(numbers)

	return
}

func swap(numbers []int, i int) {
	for j := 1; j < i; j++ {

		if numbers[j-1] > numbers[j] {
			intermediate := numbers[j]
			numbers[j] = numbers[j-1]
			numbers[j-1] = intermediate
		}
	}
	return
}
