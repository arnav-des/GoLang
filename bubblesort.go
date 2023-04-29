package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter your array (up to 10 integers): ")
	scanner.Scan()
	stringContent := scanner.Text()
	stringArray := strings.Split(stringContent, " ")
	if len(stringArray) > 10 {
		log.Fatal("Up to 10 ints allowed")
		return
	}
	integerSlice := make([]int, 0)
	for _, elem := range stringArray {
		intElem, _ := strconv.Atoi(elem)
		integerSlice = append(integerSlice, intElem)
	}
	BubbleSort(integerSlice)
	fmt.Println(integerSlice)
}

func BubbleSort(unsortedSlice []int) {
	sliceLength := len(unsortedSlice)
	for i := 0; i < sliceLength-1; i++ {
		for j := 0; j < sliceLength-i-1; j++ {
			Swap(unsortedSlice, j)
		}
	}
}

func Swap(unsortedSlice []int, counter int) {
	if unsortedSlice[counter] > unsortedSlice[counter+1] {
		temp := unsortedSlice[counter]
		unsortedSlice[counter] = unsortedSlice[counter+1]
		unsortedSlice[counter+1] = temp
	}
}
