package main

import (
	"bufio"
	"fmt"
	"os"

	"github.io/savutro/golang/services/client"
	"github.io/savutro/golang/services/server"
)

func main() {
	fmt.Println("What do you want to start? c = client, s = server, e = exit")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		text := scanner.Text()

		if text == "c" {
			client.StartClient()
		} else if text == "s" {
			server.StartServer()
		} else if text == "e" {
			break
		} else {
			fmt.Println("Invalid choice. Please enter c, s, or e.")
		}
	}
}
