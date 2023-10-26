package server

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
	"net/http"
)

var encryptionKey = "atsutczXRIS4svZH"

func receiveData(w http.ResponseWriter, r *http.Request) {
	// Read the encrypted data from the request body
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	// Decrypt the data
	decryptedData, err := decrypt(data, encryptionKey)
	if err != nil {
		http.Error(w, "Failed to decrypt data", http.StatusInternalServerError)
		return
	}

	// Print the decrypted data to the console
	fmt.Println("Received and Decrypted Data:")
	fmt.Println(string(decryptedData))

	// Respond to the client
	w.WriteHeader(http.StatusOK)
}

func decrypt(data []byte, key string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	if len(data) < aes.BlockSize {
		return nil, fmt.Errorf("data is too short to be decrypted: %v", err)
	}

	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]

	mode := cipher.NewCFBDecrypter(block, iv)
	mode.XORKeyStream(data, data)

	return data, nil
}

func StartHttpServer() {
	http.HandleFunc("/api/receiver", receiveData)
	http.ListenAndServe(":4200", nil)
	fmt.Println("Listening on :4200 ...")
}
