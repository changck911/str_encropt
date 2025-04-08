package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func EncryptAES(plaintext string, key string) (string, string, error) {
	// Create AES block cipher
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", "", fmt.Errorf("failed to create cipher: %v", err)
	}

	// 生成隨機的 IV
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return "", "", fmt.Errorf("failed to generate IV: %v", err)
	}
	// Encrypt plaintext
	ciphertext := make([]byte, len(plaintext))
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext, []byte(plaintext))

	// Convert ciphertext and IV to hex string
	return hex.EncodeToString(ciphertext), hex.EncodeToString(iv), nil
}
