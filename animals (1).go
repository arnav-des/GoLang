package main

import "fmt"

type Animal struct {
	food       string
	locomotion string
	noise      string
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
	animals := map[string]Animal{
		"cow":   {food: "grass", locomotion: "walk", noise: "moo"},
		"bird":  {food: "worms", locomotion: "fly", noise: "peep"},
		"snake": {food: "mice", locomotion: "slither", noise: "hsss"},
	}

	for {
		var animal, action string
		fmt.Print("> ")
		fmt.Scan(&animal, &action)

		if animal == "X" {
			break
		}

		switch action {
		case "eat":
			animals[animal].Eat()
		case "move":
			animals[animal].Move()
		case "speak":
			animals[animal].Speak()
		default:
			fmt.Println("Invalid action")
		}
	}
}
