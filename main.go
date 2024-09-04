package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

// EncryptedSecret represents the structure of an encrypted secret JSON.
type EncryptedSecret struct {
	Data struct {
		Password string `json:"password"`
	} `json:"data"`
}

// Function to decrypt using AES
func decryptAES(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func main() {
	// Command-line flags
	encryptionKey := flag.String("key", "", "AES encryption key")
	encodedSecret := flag.String("secret", "", "Base64-encoded secret")

	flag.Parse()

	if *encryptionKey == "" || *encodedSecret == "" {
		fmt.Println("Usage: go run main.go --key=<AES_KEY> --secret=<BASE64_ENCODED_SECRET>")
		os.Exit(1)
	}

	// Decode base64 value
	decodedSecret, err := base64.StdEncoding.DecodeString(*encodedSecret)
	if err != nil {
		log.Fatal("Failed to decode base64 secret:", err)
	}

	// Decrypt the secret using the provided AES encryption key
	decryptedSecret, err := decryptAES(decodedSecret, []byte(*encryptionKey))
	if err != nil {
		log.Fatal("Failed to decrypt secret:", err)
	}

	// Parse the decrypted secret into the struct
	var secret EncryptedSecret
	err = json.Unmarshal(decryptedSecret, &secret)
	if err != nil {
		log.Fatal("Failed to parse decrypted secret:", err)
	}

	// Output the decrypted secret
	fmt.Println("Decrypted secret:", secret.Data.Password)
}
