package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("What do you want to start? c = client, s = server, e = exit")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		text := scanner.Text()

		if text == "c" {
			startClient()
		} else if text == "s" {
			startServer()
		} else if text == "e" {
			break
		} else {
			fmt.Println("Invalid choice. Please enter c, s, or e.")
		}
	}
}
