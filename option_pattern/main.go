package main

import (
	"fmt"
	"option_pattern/house"
)

func main() {
	h := house.New(
		house.WithConcrete(),
		house.WithNumberFloors(4),
	)

	fmt.Println(h.GetNumberFloors())
}
