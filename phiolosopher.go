package main

import (
    "fmt"
    "sync"
)

const (
    numPhilosophers = 5
    numMeals       = 3
)

type Chopstick struct{ sync.Mutex }

type Philosopher struct {
    id              int
    leftChopstick   *Chopstick
    rightChopstick  *Chopstick
    mealsRemaining  int
    host            chan bool
    doneEating      chan bool
}

func (p Philosopher) eat() {
    for i := 0; i < numMeals; i++ {
        <-p.host
        p.leftChopstick.Lock()
        p.rightChopstick.Lock()

        fmt.Printf("starting to eat %d\n", p.id)
        p.mealsRemaining--
        fmt.Printf("finishing eating %d\n", p.id)

        p.rightChopstick.Unlock()
        p.leftChopstick.Unlock()
        p.host <- true
    }
    p.doneEating <- true
}

func main() {
    host := make(chan bool, 2)
    chopsticks := make([]*Chopstick, numPhilosophers)
    for i := 0; i < numPhilosophers; i++ {
        chopsticks[i] = new(Chopstick)
    }

    philosophers := make([]*Philosopher, numPhilosophers)
    for i := 0; i < numPhilosophers; i++ {
        philosophers[i] = &Philosopher{
            id:              i + 1,
            leftChopstick:   chopsticks[i],
            rightChopstick:  chopsticks[(i+1)%numPhilosophers],
            mealsRemaining:  numMeals,
            host:            host,
            doneEating:      make(chan bool),
        }
        go philosophers[i].eat()
    }

    // allow 2 philosophers to eat at a time
    host <- true
    host <- true

    // wait for all philosophers to finish eating
    for _, p := range philosophers {
        <-p.doneEating
    }
}
