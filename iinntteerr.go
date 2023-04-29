package main

import (
	"fmt"
	"strings"
)

type animal interface {
	Eat()
	Move()
	Noise()
}

//Cow type
type Cow struct{}

//Eat method
func (c *Cow) Eat() { fmt.Print("Grass") }

//Move method
func (c *Cow) Move() { fmt.Print("Walk") }

//Noise method
func (c *Cow) Noise() { fmt.Print("Moo") }

//Bird type
type Bird struct{}

//Eat method
func (c *Bird) Eat() { fmt.Print("Worm") }

//Move method
func (c *Bird) Move() { fmt.Print("Fly") }

//Noise method
func (c *Bird) Noise() { fmt.Print("Peep") }

//Snake type
type Snake struct{}

//Eat method
func (c *Snake) Eat() { fmt.Print("Mice") }

//Move method
func (c *Snake) Move() { fmt.Print("Crawl") }

//Noise method
func (c *Snake) Noise() { fmt.Print("Hsss") }

func main() {

	var first string
	var second string
	var third string
	var animals = make(map[string]animal)

	for {
		first = ""
		second = ""
		third = ""
		fmt.Print("\n> ")
		fmt.Scan(&first)
		fmt.Scan(&second)
		fmt.Scan(&third)

		if first == "newanimal" && second != "" && third != "" {
			third = strings.ToLower(third)

			switch {
			case third == "cow":
				animals[second] = &Cow{}
			case third == "bird":
				animals[second] = &Bird{}
			case third == "snake":
				animals[second] = &Snake{}
			}

			fmt.Println("Created it!")
		}

		if first == "query" && second != "" && third != "" {
			third = strings.ToLower(third)

			switch {
			case third == "eat":
				animals[second].Eat()
			case third == "move":
				animals[second].Move()
			case third == "noise":
				animals[second].Noise()
			}
		}
	}
}
