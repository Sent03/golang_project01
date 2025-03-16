package main

import "fmt"

type Engine struct {
	Power int
}

type Wheels struct {
	Count int
}

type Car struct {
	Engine Engine
	Wheels Wheels
}

func (c *Car) Drive() string {
	return fmt.Sprintf("Автомобиль едет с двигателем мощностью %d л.с. и %d колесами!", c.Engine.Power, c.Wheels.Count)
}

func main() {

}
