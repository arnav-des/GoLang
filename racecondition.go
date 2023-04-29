package main

import (
	"fmt"
	"time"
)

func execute(some string) {

	for i := 1; true; i++ {

		fmt.Println(some, i)
		
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {

	go execute("First")
	
	go execute("Second")
	
	fmt.Println("Program ends successfully")
	
}
