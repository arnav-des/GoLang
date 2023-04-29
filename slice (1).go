package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var s []int = make([]int, 3)
	var in string
	fmt.Println("Enter an interger or press X to exit:")
	for true {
		fmt.Scanln(&in)
		if in == "X" {
			break
		}
		ap, err := strconv.Atoi(in)
		if err != nil {
			fmt.Println("Wrong input")
			continue
		}
		s = append(s, ap)
		sort.Ints(s[:])
		fmt.Println(s)
		fmt.Println("Enter an interger or press X to exit::")
	}
}
