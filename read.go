package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("names.txt") //open the file
	if err != nil {
		fmt.Println("Error opening file: ", err)
		os.Exit(1)
	}
	defer f.Close() //close file at the end of the program

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		name := scanner.Text()
		names := strings.Split(name, " ") //split the name into first and last name
		fmt.Printf("First Name: %s, Last Name: %s\n", names[0], names[1]) //print the first and last name
	}

}