package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type funcType func(float64) float64

func GenDisplaceFn(a float64, v float64, s float64) funcType {
	return func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v*t + s
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var fn funcType

	fmt.Print("Enter Acceleration: ")
	scanner.Scan()
	a, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Print("Enter Initial Velocity: ")
	scanner.Scan()
	v, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Print("Enter Initial Displacement: ")
	scanner.Scan()
	s, _ := strconv.ParseFloat(scanner.Text(), 64)

	fn = GenDisplaceFn(a, v, s)

	for {
		fmt.Print("Enter time value(\"quit\" to exit): ")
		scanner.Scan()
		input := scanner.Text()
		if strings.ToLower(input) == "quit" {
			break
		}
		time, err := strconv.ParseFloat(input, 64)
		if err != nil {
			println("Invalid value")
			continue
		}
		fmt.Printf("Displacement after time %.2f is %.2f\n", time, fn(time))
	}
}
