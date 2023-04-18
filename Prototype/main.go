package main

import "fmt"

type Animal interface {
	SaySomething()
	Clone() Animal
}

type Cat struct {
	Name  string
	Age   int
	Sound string
}

func (c *Cat) Clone() Animal {
	return &Cat{
		Name:  c.Name,
		Age:   c.Age,
		Sound: c.Sound,
	}
}

func (c *Cat) SaySomething() {
	fmt.Printf("%d yeard old %s makes %s", c.Age, c.Name, c.Sound)
}

func main() {
	var a1, a2 Animal

	a1 = &Cat{
		Name:  "Kit",
		Age:   10,
		Sound: "Meow",
	}

	a2 = a1.Clone()

	a2.(*Cat).Sound = "ShhHEHheh"
	a2.(*Cat).Age = 18

	a1.SaySomething()
	a2.SaySomething()
}
