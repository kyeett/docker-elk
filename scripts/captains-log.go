package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type Signal struct {
	to      string
	from    string
	message string
	testId  string
}

func generateSignal() Signal {

	names := []string{
		"Alice",
		"Bjorn",
		"Cilla",
		"Dylan",
	}
	to := names[rand.Intn(len(names))]
	from := names[rand.Intn(len(names))]

	messages := []string{
		"Hello!",
		"Who are you?",
		"I am " + from,
		"My name is " + from,
		"Goodbye",
	}

	message := Signal{
		to:      to,
		from:    from,
		message: messages[rand.Intn(len(messages))],
		testId:  "DSA",
	}

	return message
}

func main() {
	rand.Seed(time.Now().Unix())
	fmt.Println("Captain's log:\n")
	signal := generateSignal()
	b, err := json.Marshal(signal)

	if err != nil {
		fmt.Println("Error encoding:", err.Error())
	}

	fmt.Println(signal)
	fmt.Println(b)
}
