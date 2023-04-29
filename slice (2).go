package main

import (
	"fmt"
	"strconv"
)

func main() {
	var user_input string
	sli := make([]int, 0, 3)

	for i:=0; i<len(sli)+1; i++ {
		fmt.Print("Enter a Number or X: ")
		fmt.Scan(&user_input)

		//if found X break the loop
		if user_input == "X"{
			fmt.Print	("User has decided to quit, thank you! ")
			break
		}

		num, err := strconv.Atoi(user_input)
		if err != nil {
			fmt.Println(err)
		} else {
			sli = append(sli, num)
			sortslice(sli)
			fmt.Println("The sorted slice is::")
			fmt.Println(sli)
		}
	}
}

//sort the slice
func sortslice(slc []int) []int{
	var n int
	var sorted bool
	n = len(slc)
	sorted = false

	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if slc[i] > slc[i+1] {
				slc[i+1], slc[i] = slc[i], slc[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
	return slc
}
