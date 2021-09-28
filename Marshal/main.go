package main

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	Nombre   string
	Apellido string
	Edad     int
}

func main() {

	p1 := persona{
		Nombre:   "Alejandro",
		Apellido: "Carballo",
		Edad:     41,
	}

	p2 := persona{
		Nombre:   "Melisa",
		Apellido: "Aime",
		Edad:     35,
	}

	personas := []persona{p1, p2}

	fmt.Println(personas)

	bs, err := json.Marshal(personas)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
