package main
import (
   "fmt"
   "sort"
)
func main(){
   fmt.Printf("Enter size of your array: ")
   var size int
   fmt.Scanln(&size)
   var arr = make([]int, size)
   for i:=0; i<size; i++ {
      fmt.Printf("Enter %dth element: ", i)
      fmt.Scanf("%d", &arr[i])
   }
   fmt.Println("Your array is: ", arr)
   sort.Slice(arr, func(i, j int) bool {
    return arr[i] < arr[j]

	for _, v := range arr {
    	fmt.Println(v)
	}
})
}