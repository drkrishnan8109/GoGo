package main

import (
	"bufio"
	"fmt"
	"net"
)

type response string

func handleConnection(conn net.Conn) {
	defer conn.Close()
	msg, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Printf("msg: %v\n", msg)
	if err != nil {
		fmt.Println("Error reading from connection", err)
		return
	}
	_, err = conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("Error writing to connection:", err)
	}
	fmt.Println("send my echo!", msg)
	return
}

func main() {
	listener, err := net.Listen("tcp", ": 8080")
	if err != nil {
		// Use nc -vz localhost 8080 to test connection to the server
		fmt.Println("Error listening to tcp port:8080", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening on port 8080...")

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Error accepting connection", err)
			continue
		}
		go handleConnection(conn)
	}
}
