package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
  arr := arrayInput()
  n := len(arr)
  n1, n2, n3 := n/4, 2*n/4, 3*n/4 

  c := make(chan []int, 3)
  
  arr1 := arr[:n1]
  arr2 := arr[n1: n2]
  arr3 := arr[n2: n3]
  arr4 := arr[n3:]

  go mergeSort(arr1, arr2, c)
  go mergeSort(arr3, arr4, c)
  first := <- c
  second := <- c

  go mergeSort(first, second, c)
  sortedSlice := <- c

  fmt.Println("The sorted array is - ", sortedSlice)
}

func mergeSort(a1, a2 []int, c chan []int) {
  fmt.Println(a1, a2)
  result := []int{}

  for len(a1) > 0 && len(a2) > 0 {
    if a1[0] <= a2[0] {
      result = append(result, a1[0])
      a1 = a1[1:]
    } else {
      result = append(result, a2[0])
      a2 = a2[1:]
    }
  }

  for len(a1) > 0 {
    result = append(result, a1[0])
    a1 = a1[1:]
  }
  for len(a2) > 0 {
    result = append(result, a2[0])
    a2 = a2[1:]
  }

  c <- result
}

// func arrayInput() []int {
//     var n int
//     fmt.Scan(&n)
//     arr := make([]int, n)
//     for i := 0; i < n; i++ {
//         fmt.Scan(&arr[i])
//     }
//     return arr
// }

func arrayInput() []int {
  scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter space separated numbers: ")
	if !scanner.Scan() {
		err := scanner.Err()
		if err == nil {
			err = errors.New("no input received")
		}
		fmt.Printf("Error: %v\n", err)
		return []int{}
	}

	input := scanner.Text()
	numStrs := strings.Split(input, " ")
	nums := make([]int, len(numStrs))
	for i, numStr := range numStrs {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Printf("Error: invalid number %q at position %d\n", numStr, i)
			return []int{}
		}
		nums[i] = num
	}
  return nums
}
