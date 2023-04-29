package main

import (
	"fmt"
	"strconv"
)

func GenDisplaceFn(acc, vo, so float64) func(t float64) float64 {
	return func(t float64) float64 {
		return 0.5*acc*t*t + vo*t + so
	}
}

func GetInitialConditions() (float64, float64, float64) {
	cmd_inputs := make([]string, 3)
	_, _ = fmt.Scan(&cmd_inputs[0], &cmd_inputs[1], &cmd_inputs[2])
	var outputs [3]float64 = [3]float64{0, 0, 0}
	for i, element := range cmd_inputs {
		outputs[i], _ = strconv.ParseFloat(element, 64)
	}
	return outputs[0], outputs[1], outputs[2]
}

func GetSingleInput() float64 {
	var cmd_input string
	_, _ = fmt.Scan(&cmd_input)
	return_val, _ := strconv.ParseFloat(cmd_input, 64)
	return return_val
}

func main() {
	fmt.Print("Etner acceleration, initial velcoity and initial displacement as space-separated values: ")
	acc, v0, s0 := GetInitialConditions()
	displace_func := GenDisplaceFn(acc, v0, s0)
	fmt.Print("\nEnter the time as a single value: ")
	time := GetSingleInput()
	fmt.Print("\nThe final displacement is ")
	fmt.Print(displace_func(time))
}
