package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

func EncryptAES(plaintext string, iv string, key string) (string, error) {
	// Create AES block cipher
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %v", err)
	}

	// Encrypt plaintext
	ciphertext := make([]byte, len(plaintext))
	stream := cipher.NewCFBEncrypter(block, []byte(iv))
	stream.XORKeyStream(ciphertext, []byte(plaintext))

	// Convert ciphertext to hex string
	return hex.EncodeToString(ciphertext), nil
}
