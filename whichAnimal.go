package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Animal struct
type Animal struct {
	food       string
	locomotion string
	noise      string
}

var cow = Animal{food: "grass", locomotion: "walk", noise: "moo"}
var bird = Animal{food: "worms", locomotion: "fly", noise: "peep"}
var snake = Animal{food: "mice", locomotion: "slither", noise: "hsss"}

//Eat returns what animal eats
func (animal Animal) Eat() string {
	return animal.food
}

//Move returns what animal eats
func (animal Animal) Move() string {
	return animal.locomotion
}

//Speak returns what animal eats
func (animal Animal) Speak() string {
	return animal.noise
}

func promptUser() (string, string) {
	fmt.Print(">")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userInput := scanner.Text()
	userinputArray := strings.Fields(userInput)
	return userinputArray[0], userinputArray[1]
}
func main() {
	fmt.Println("Enter animal name and information requested by a single space:")
	for {
		animal, information := promptUser()
		switch animal {
		case "cow":
			switch information {
			case "eat":
				fmt.Println(cow.Eat())
			case "move":
				fmt.Println(cow.Move())
			case "speak":
				fmt.Println(cow.Speak())
			default:
				fmt.Println("Wrong input for first field!")
			}

		case "bird":
			switch information {
			case "eat":
				fmt.Println(bird.Eat())
			case "move":
				fmt.Println(bird.Move())
			case "speak":
				fmt.Println(bird.Speak())
			default:
				fmt.Println("Wrong input for first field!")
			}

		case "snake":
			switch information {
			case "eat":
				fmt.Println(snake.Eat())
			case "move":
				fmt.Println(snake.Move())
			case "speak":
				fmt.Println(snake.Speak())
			default:
				fmt.Println("Wrong input for first field!")
			}
		default:
			fmt.Println("Wrong input for second input!")
		}
	}

}
