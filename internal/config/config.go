package config

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func ReadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file")
	}

	method := os.Getenv("METHOD")
	key := os.Getenv("AES_KEY_STR")
	iv := os.Getenv("AES_IV_STR")

	if method == "" || key == "" || iv == "" || os.Getenv("STR") == "" {
		return fmt.Errorf("METHOD is not set in .env file")
	}

	os.Setenv("AES_KEY", compressStringTo16Chars(key))
	os.Setenv("AES_IV", compressStringTo16Chars(iv))

	return nil
}

func compressStringTo16Chars(input string) string {
	// 計算 MD5
	hash := md5.Sum([]byte(input))

	// 取前 16 個字元作為 AES 密鑰
	str := hex.EncodeToString(hash[:])[:16]

	return str
}
