package main
import "fmt"

type Vehicle struct {
	name string
}

func (v *Vehicle) Move() string {
	return v.name + " " + "едет"
}

type Car struct {
	Vehicle Vehicle
}

func (c *Car) Move() string {
	return c.Vehicle.name + " " + "едет"
}

type Bicycle struct {
	Vehicle Vehicle
}

func (b *Bicycle) Move() string {
	return b.Vehicle.name + " " + "едет"
}

func main {
	
}