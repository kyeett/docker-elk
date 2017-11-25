package main

import (
	"encoding/json"
	"fmt"
)

type Signal struct {
	to   string
	from string
}

func generateSignal() Signal {
	message := Signal{
		to:   "Magnus",
		from: "Bjorn",
	}

	return message
}

func main() {
	fmt.Println("Captain's log:\n")
	signal := generateSignal()
	b, err := json.Marshal(signal)

	if err != nil {
		fmt.Println("Error encoding:", err.Error())
	}

	fmt.Println(signal)
	fmt.Println(b)
}
