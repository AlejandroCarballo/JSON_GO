package main

import (
	"encoding/json"
	"fmt"
)

type usuario struct {
	Name string
	Edad int
}

func main() {
	u1 := usuario{
		Name: "Alejandro",
		Edad: 42,
	}
	u2 := usuario{
		Name: "Meli",
		Edad: 35,
	}
	u3 := usuario{
		Name: "Camilo",
		Edad: 4,
	}

	usuarios := []usuario{u1, u2, u3}
	fmt.Println(usuarios)

	bs, err := json.Marshal(usuarios)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}
