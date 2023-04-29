package main

import "fmt"

// Define an Animal interface type
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Define specific types of animals
type Cow struct {
	name string
}

type Bird struct {
	name string
}

type Snake struct {
	name string
}

// Define methods for each type of animal
func (c Cow) Eat() {
	fmt.Println(c.name, "eats grass")
}

func (c Cow) Move() {
	fmt.Println(c.name, "walks")
}

func (c Cow) Speak() {
	fmt.Println(c.name, "says moo")
}

func (b Bird) Eat() {
	fmt.Println(b.name, "eats worms")
}

func (b Bird) Move() {
	fmt.Println(b.name, "flies")
}

func (b Bird) Speak() {
	fmt.Println(b.name, "says chirp")
}

func (s Snake) Eat() {
	fmt.Println(s.name, "eats mice")
}

func (s Snake) Move() {
	fmt.Println(s.name, "slithers")
}

func (s Snake) Speak() {
	fmt.Println(s.name, "hisses")
}

func main() {
	fmt.Println(">")
	var command string
	var animal Animal
	for {
		fmt.Scan(&command)
		if command == "newanimal" {
			var name, animalType string
			fmt.Scan(&name, &animalType)
			if animalType == "cow" {
				animal = Cow{name}
			} else if animalType == "bird" {
				animal = Bird{name}
			} else if animalType == "snake" {
				animal = Snake{name}
			}
			fmt.Println("Created it!")
		} else if command == "query" {
			var info string
			fmt.Scan(&info)
			if info == "eat" {
				animal.Eat()
			} else if info == "move" {
				animal.Move()
			} else if info == "speak" {
				animal.Speak()
			}
		}
		fmt.Println(">")
	}
}