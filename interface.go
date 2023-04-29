package main

import(
	"fmt"
	"strings"
	"bufio"
	"os"
)

type animal interface {
    speak()
	eat()
	move()
}

type Cow struct{
	food string
	sound string
	movement string
}

type Bird struct{
	food string
	sound string
	movement string
}

type Snake struct{
	food string
	sound string
	movement string
}

func (cow Cow) eat(){
	fmt.Printf("Eat:grass\n")
}

func (cow Cow) speak(){
	fmt.Printf("Speak:moo\n")
}
func (cow Cow) move(){
	fmt.Printf("Move:walk\n")
}

func (bird Bird) eat(){
	fmt.Printf("Eat:worms\n")
}

func (bird Bird) speak(){
	fmt.Printf("Speak:peep\n")
}
func (bird Bird) move(){
	fmt.Printf("Move:fly\n")
}


func (snake Snake) eat(){
	fmt.Printf("Eat:mice\n")
}

func (snake Snake) speak(){
	fmt.Printf("Speak:hsss\n")
}
func (snake Snake) move(){
	fmt.Printf("Move:slither\n")
}


func main(){
	reader := bufio.NewReader(os.Stdin)
	cow1 := make(map[string]Cow)
	snake1 := make(map[string]Snake)
	bird1 := make(map[string]Bird)
	var input []string
	fmt.Println("Enter x to esacpe :)") 
	for{
		fmt.Print(">")
		input1,_ := reader.ReadString('\n')
		input1 = strings.TrimRight(input1, "\r\n")
		input1 = strings.ToLower(input1)
		input = strings.Split(input1, " ")
		if (input[0] == "x"){
			break
		}else if (input[0] == "newanimal"){
				if(input[2] == "cow"){
					name1 := input[1]
					name3 := Cow{"grass","moo","walk"}
					cow1[name1] = name3
					fmt.Println("Created it!")
				}else if(input[2] == "bird"){
					name1 := input[1]
					name3 := Bird{"worms","peep","fly"}
					bird1[name1] = name3
					fmt.Println("Created it!")
				}else if(input[2] == "snake"){
					name1 := input[1]
					name3 := Snake{"mice","hsss","slither"}
					snake1[name1] = name3
					fmt.Println("Created it!")
				}else{
					fmt.Println("Wrong Input")
				}
			}else if (input[0] == "query"){
				if(input[2] == "speak"){
					name2 := input[1]
					if v, found := cow1[name2]; found {
						v.speak()
					}
					if v, found := bird1[name2]; found {
						v.speak()
					}
					if v, found := snake1[name2]; found {
						v.speak()
					}
				}else if(input[2] == "move"){
					name2 := input[1]
					if v, found := cow1[name2]; found {
						v.move()
					}
					if v, found := bird1[name2]; found {
						v.move()
					}
					if v, found := snake1[name2]; found {
						v.move()
					}
			}else if(input[2] == "eat"){
				name2 := input[1]
				if v, found := cow1[name2]; found {
					v.eat()
				}
				if v, found := bird1[name2]; found {
					v.eat()
				}
				if v, found := snake1[name2]; found {
					v.eat()
				}
		}
	}else{
		fmt.Println("Wrong Input")
	}
}
}