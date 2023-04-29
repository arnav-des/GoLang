package main

import (
    "fmt"
)

func GenDisplaceFn(a,vo,so float64) func(float64) float64 {
    return func(t float64) float64 {
        return 0.5*a*t*t + vo*t + so
    }
}

func main() {
    var acceleration, initialVelocity, initialDisplacement, time float64
    fmt.Println("Please enter values of acceleration, initial velocity and initial displacement:")
    fmt.Scan(&acceleration, &initialVelocity, &initialDisplacement)

    displacementFn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

    fmt.Println("Please enter value for time:")
    fmt.Scan(&time)

    fmt.Println("Displacement after entered time is:",displacementFn(time))
}