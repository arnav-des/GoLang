package main

import "fmt"

func BubbleSort(slice []int) {
    for i := 0; i < len(slice); i++ {
        for j := 0; j < len(slice)-1-i; j++ {
            if slice[j] > slice[j+1] {
                swap(slice, j, j+1)
            }
        }
    }
}

func swap(slice []int, i, j int) {
    temp := slice[i]
    slice[i] = slice[j]
    slice[j] = temp
}

func main() {
    arr := []int{8, 7, 4, 3, 2, 1, 0}
    fmt.Println("Unsorted Array: ", arr)

    BubbleSort(arr)
    fmt.Println("Sorted Array: ", arr)
}