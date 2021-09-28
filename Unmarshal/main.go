package main

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	Nombre   string `json:"Nombre"`
	Apellido string `json:"Apellido"`
	Edad     int    `json:"Edad"`
}

func main() {
	s := `[{"Nombre":"Alejandro","Apellido":"Carballo","Edad":41},{"Nombre":"Melisa","Apellido":"Aime","Edad":35}]`
	bs := []byte(s)

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var personas []persona
	err := json.Unmarshal(bs, &personas)
	if err != nil {
		fmt.Println(err)
	}

	for i, v := range personas {
		fmt.Println("\nPersona Numero", i+1)
		fmt.Println(v.Nombre, v.Apellido, v.Edad)

	}

}
