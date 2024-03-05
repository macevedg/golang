package main

import (
	"errors"
	"fmt"
)

func main() {

	var param []Human

	student := Student{
		age:  10,
		Name: "Fabri",
	}

	worker := Worker{
		age:          40,
		Name:         "Andres",
		SocialNumber: 71795579,
	}

	param = append(param, &student, &worker)

	printer := Printer{}
	printer.printNames(param)

	figher := KameSenin{
		fighter: &worker,
	}

	fmt.Println("FIGHTER NAME:", figher.fighter.ShowName())
}

func (p *Printer) printNames(humans []Human) (string, error) {

	if len(humans) == 0 {
		return "", errors.New("not available humans")
	}
	names := ""
	for _, h := range humans {
		names += h.ShowName()
		fmt.Println(h.ShowName())
	}

	return names, nil
}

type Printer struct {
	human []Human
}

func printNames(h Human) {
	fmt.Println(h.ShowName())
}

type Human interface {
	ShowName() string
}

type Student struct {
	Name string
	age  int
}

func (s *Student) ShowName() string {
	return s.Name
}

type Worker struct {
	Name         string
	age          int
	SocialNumber int
}

func (w *Worker) ShowName() string {
	return w.Name
}

type KameSenin struct {
	fighter Human
}
