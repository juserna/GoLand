package main

import "fmt"

// Interfaces
// Las interfaces son implicitas
// struct -> interface

func main() {
	fmt.Println("I want a Coffee")

	var (
		i ItalianCoffeeMachine
		c ColombianCoffeeMachine
	)

	italianCofee := GetCoffee(&i, 10)
	italianCofee.PrintCoffee()

	colombianCoffee := GetCoffee(&c, 19)
	colombianCoffee.PrintCoffee()

}

type Coffee struct {
	Intensity int
	Region    string
}

// CoffeeMaker
type CoffeeMaker interface {
	MakeCoffee(intensity int) Coffee
}

func (c *Coffee) PrintCoffee() {
	fmt.Println(fmt.Sprintf("This coffee is from %s and intensity is %d", c.Region, c.Intensity))
}

type ItalianCoffeeMachine struct {
}

type ColombianCoffeeMachine struct {
}

func (i *ItalianCoffeeMachine) MakeCoffee(intensity int) Coffee {
	return Coffee{Intensity: intensity, Region: "Italy"}
}

func (c *ColombianCoffeeMachine) MakeCoffee(intensity int) Coffee {
	return Coffee{Intensity: intensity, Region: "Colombia"}
}

func GetCoffee(coffeeMaker CoffeeMaker, i int) Coffee {
	return coffeeMaker.MakeCoffee(i)
}
