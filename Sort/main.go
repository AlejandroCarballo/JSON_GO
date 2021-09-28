package main

import (
	"fmt"
	"sort"
)

func main() {

	x1 := []int{4, 7, 32, 99, 18, 16, 56, 12}
	x2 := []string{"Mi nombre", "es niña y soy la bianca", " la casa mas hermosa de .."}
	sort.Strings(x2)
	fmt.Println(x2)
	sort.Ints(x1)
	fmt.Println(x1)

}
