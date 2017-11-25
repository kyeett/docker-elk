package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

type Signal struct {
	To      string
	From    string
	Message string
	TestId  string
}

func generateSignal() Signal {

	names := []string{
		"Alice",
		"Bjorn",
		"Cilla",
		"Dylan",
	}

	// Randomize to and from names
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
		To:      to,
		From:    from,
		Message: messages[rand.Intn(len(messages))],
		TestId:  "DSA",
	}

	return message
}

func handleRequest(conn net.Conn) {
	signal := generateSignal()
	b, err := json.Marshal(signal)

	if err != nil {
		fmt.Println("Error encoding:", err.Error())
	}

	fmt.Println(signal)
	fmt.Println(b)

	conn.Write(b)
	conn.Close()
}

func main() {
	rand.Seed(time.Now().Unix())

	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting: ", err.Error())
		os.Exit(1)
	}
	handleRequest(conn)

}
