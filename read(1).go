package main

import (
	"bufio"
	"fmt"
	"os"
)

type Name struct {
	fname [20]byte
	lname [20]byte
}

func main() {

	var names []Name

	fmt.Print("Enter the name of file:\n")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	filename := scanner.Text()

	file, err := os.Open(filename)

	if err != nil {
		fmt.Print("Error opening the file", err)
		return
	}

	defer file.Close()

	scanner = bufio.NewScanner(file)

	var n Name

	for scanner.Scan() {

		fmt.Sscanf(scanner.Text(), "%s %s", &n.fname, &n.lname)
		names = append(names, n)
	}

	if err := scanner.Err(); err != nil {
		fmt.Print("Error reading the file: ", err)
		return
	}

	fmt.Print("\nnames: ")
	for _, n = range names {
		fmt.Printf("First Name: %s, Last Name: %s", n.fname, n.lname)
	}

}
