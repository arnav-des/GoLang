package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food string
	locomotion string
	noise string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}
	for {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input := scanner.Text()
			inputSlices := strings.Split(input, " ")
			animalName := inputSlices[0]
			animalInfo := inputSlices[1]

			switch animalName {
			case "cow":
				switch animalInfo {
				case "eat":
					cow.Eat()
				case "move":
					cow.Move()
				case "speak":
					cow.Speak()
				}
			case "bird":
				switch animalInfo {
				case "eat":
					bird.Eat()
				case "move":
					bird.Move()
				case "speak":
					bird.Speak()
				}
			case "snake":
				switch animalInfo {
				case "eat":
					snake.Eat()
				case "move":
					snake.Move()
				case "speak":
					snake.Speak()
				}
			}
		}
	}
}