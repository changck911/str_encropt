package config

import (
	"crypto/rand"
	"fmt"
	"os"

	"crypto/sha256"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/pbkdf2"
)

func ReadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file")
	}

	method := os.Getenv("METHOD")
	key := os.Getenv("AES_KEY_STR")
	salt := os.Getenv("AES_SALT")

	if method == "" || key == "" || os.Getenv("STR") == "" {
		return fmt.Errorf("METHOD is not set in .env file")
	}

	// 使用 PBKDF2 生成安全的密鑰
	if salt == "" {
		salt = "default_salt_please_change_in_production"
	}
	derivedKey := deriveKey(key, salt)
	os.Setenv("AES_KEY", string(derivedKey[:16])) // 取前16字節作為AES密鑰

	// 生成隨機的 IV
	iv := generateRandomIV()
	os.Setenv("AES_IV", iv)

	return nil
}

func deriveKey(password, salt string) []byte {
	// 使用 PBKDF2 與 SHA-256，進行 10000 次迭代，生成 32 字節的密鑰
	return pbkdf2.Key([]byte(password), []byte(salt), 10000, 32, sha256.New)
}

// 生成隨機的 16 字節 IV
func generateRandomIV() string {
	iv := make([]byte, 16)
	if _, err := rand.Read(iv); err != nil {
		// 如果無法生成隨機IV，使用默認值
		return "0123456789abcdef"
	}
	return string(iv)
}
