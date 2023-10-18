package server

import (
	"fmt"
	"net"
	"os"
)

func StartServer() {
	// Listen on a specific port for incoming connections
	listener, err := net.Listen("tcp", "localhost:12345")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 12345...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Create a buffer to hold the incoming data
	buffer := make([]byte, 1024)

	// Create a file to save the received data
	file, err := os.Create("received_file.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	fmt.Println("Receiving a file...")

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading data:", err)
			break
		}

		_, err = file.Write(buffer[:n])
		if err != nil {
			fmt.Println("Error writing data to file:", err)
			break
		}
	}

	fmt.Println("File received and saved as 'received_file.txt'")
}
