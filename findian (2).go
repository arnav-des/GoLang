package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
	var inp  string
	var inp2 string
	cond1, cond2, cond3 := false, false, false

	fmt.Printf("Enter a string: ")
	//fmt.Scan(&inp)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inp = scanner.Text()

	inp2 = strings.Trim(strings.ToUpper(inp), " ")
	
	for i := 0; i< len(inp2); i++ {
		// fmt.Printf("%c\n", inp2[i])
		c := inp2[i]
		if i == 0 && c == 'I' {
			cond1 = true
		}
		if i > 0 && c == 'A' {
			cond2 = true
		}
		if i == len(inp2) - 1 && c == 'N' {
			cond3 = true
		}
	}
	if cond1 && cond2 && cond3 {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not Found!\n")
	}
}
