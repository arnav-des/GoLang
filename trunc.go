package main
import ("fmt")

func main() {
  var i float32

  fmt.Print("Type a number: ")
  fmt.Scan(&i)
  fmt.Println("Your number without truncation is:", i)
  fmt.Printf("Your number with truncation is: %.3f\n", i)
}