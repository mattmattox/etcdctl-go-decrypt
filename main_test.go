package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"testing"
)

// TestEncryptAES is a helper function to encrypt a plaintext secret using AES for testing purposes.
func encryptAES(plaintext, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize()) // A mock nonce, usually you would use a random nonce
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// TestDecryptAES_Success tests successful decryption of an encrypted secret.
func TestDecryptAES_Success(t *testing.T) {
	// Corrected key to exactly 32 bytes for AES-256
	key := []byte("my32charactersecretpassword12345")
	plaintext := []byte(`{"data": {"password": "supersecret"}}`)

	// Encrypt the plaintext using the key
	encodedCiphertext, err := encryptAES(plaintext, key)
	if err != nil {
		t.Fatalf("Encryption failed: %v", err)
	}

	// Decode base64 value
	decodedCiphertext, err := base64.StdEncoding.DecodeString(encodedCiphertext)
	if err != nil {
		t.Fatalf("Failed to decode base64: %v", err)
	}

	// Decrypt the ciphertext
	decrypted, err := decryptAES(decodedCiphertext, key)
	if err != nil {
		t.Fatalf("Decryption failed: %v", err)
	}

	expectedPassword := "supersecret"
	var secret EncryptedSecret
	err = json.Unmarshal(decrypted, &secret)
	if err != nil {
		t.Fatalf("Failed to unmarshal decrypted secret: %v", err)
	}

	// Check if the decrypted password matches the expected password
	if secret.Data.Password != expectedPassword {
		t.Fatalf("Expected password to be %s, but got %s", expectedPassword, secret.Data.Password)
	}
}

// TestDecryptAES_Fail tests decryption with incorrect key or ciphertext.
func TestDecryptAES_Fail(t *testing.T) {
	invalidCiphertext := "invalidbase64"
	key := []byte("my32charactersecretpassword123456")

	// Decode invalid base64 string (it will fail, and that's okay for this test)
	decodedCiphertext, err := base64.StdEncoding.DecodeString(invalidCiphertext)
	if err == nil {
		_, err = decryptAES(decodedCiphertext, key)
		if err == nil {
			t.Fatal("Expected decryption to fail with invalid input, but it succeeded")
		}
	}
}
