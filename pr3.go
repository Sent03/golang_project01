package main
import "fmt"

type Drawable interface {
	Draw() string
}

type Circle struct {
	name string
}

func (c *Circle) Draw() string {
	return "Рисует" + " " + c.name
}

type Square struct {
	name string
}

func (s *Square) Draw() string {
	return "Рисует" + " " + s.name
}


func main {
	
}