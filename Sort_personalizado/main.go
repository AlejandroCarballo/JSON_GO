package main

import (
	"fmt"
	"sort"
)

type Persona struct {
	Nombre string
	Edad   int
}

type PorEdad []Persona

func (a PorEdad) Len() int {
	return len(a)
}
func (a PorEdad) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a PorEdad) Less(i, j int) bool {
	return a[i].Edad < a[j].Edad
}

func main() {

	p1 := Persona{"Alejandro", 42}
	p2 := Persona{"Melisa", 35}
	p3 := Persona{"Bianca", 8}
	p4 := Persona{"Camilo", 4}

	personas := []Persona{p1, p2, p3, p4}

	sort.Sort(PorEdad(personas))
	fmt.Println(PorEdad(personas))
}
