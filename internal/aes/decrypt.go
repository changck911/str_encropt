package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

func DecryptAES(ciphertextHex string, iv string, key string) (string, error) {
	// Decode ciphertext from hex string
	ciphertext, err := hex.DecodeString(ciphertextHex)
	if err != nil {
		return "", fmt.Errorf("invalid ciphertext: %v", err)
	}

	// Create AES block cipher
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %v", err)
	}

	// Decrypt ciphertext
	plaintext := make([]byte, len(ciphertext))
	stream := cipher.NewCFBDecrypter(block, []byte(iv))
	stream.XORKeyStream(plaintext, ciphertext)

	return string(plaintext), nil
}
