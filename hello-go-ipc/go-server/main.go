package main // import "go-server"

import (
	// "bufio"

	"bufio"
	"fmt"
	"os"

	"namedpipe"
)

// Message - 메시지
type Message struct {
	Message string
}

func acceptPipeConnections(pipe namedpipe.NamedPipeServer) {
	reqID := "dummyreq"
	fmt.Println("Listening for pipe connections...")
	err := pipe.NewClient(reqID)
	if err != nil {
		fmt.Printf("Error: %+v", err)
		return
	}

	fmt.Println("New client connected...")
	/*
		Uncomment below lines to send messages to client

		// message := Message{
		// 	Message: "Hello, World from server !",
		// }
		// writer, err := pipe.GetWriter(reqID)
		// if err != nil {
		// 	fmt.Printf("Unable to get writer: %+v\n", err)
		// 	return
		// }
		// enc := gob.NewEncoder(writer)
		// err = enc.Encode(message)
		// if err != nil {
		// 	fmt.Printf("Unable to write: %+v\n", err)
		// 	return
		// }

		// fmt.Println("Sent message to client...")
	*/
	for {
		reader, err := pipe.GetReader(reqID)
		if err != nil {
			fmt.Printf("Unable to get reader: %+v", err)
			return
		}
		fmt.Println("got reader...")

		r := bufio.NewReader(reader)
		msg, err := r.ReadString('\n')
		if err != nil {
			// handle error
			fmt.Printf("error 1: %s\n", err)
			return
		}
		fmt.Printf("Received response: %+v", msg)
	}
}

// CreateServer - 서버 시작
func CreateServer() {
	fmt.Println("Attempting to create pipe")
	pipeName := "helloNamedPipe"
	pipe, err := namedpipe.NewNamedPipeServer(pipeName)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
		os.Exit(1)
	}
	fmt.Println("Created pipe")
	acceptPipeConnections(pipe)
}

func main() {
	CreateServer()
}
