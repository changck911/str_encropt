package internal

import (
	"fmt"
	"main/internal/aes"
	"main/internal/config"
	"os"
	"strings"
)

func New() {
	if err := config.ReadEnv(); err != nil {
		panic(err.Error())
	}

	switch os.Getenv("METHOD") {
	case "1":
		Encrypt()
	case "2":
		Decrypt()
	default:
		panic("METHOD ERROR!!!")
	}
}

func Encrypt() {
	ciphertext, iv, err := aes.EncryptAES(os.Getenv("STR"), os.Getenv("AES_KEY"))
	if err != nil {
		panic(err.Error())
	}

	// 將加密結果和IV一起輸出，用冒號分隔
	result := fmt.Sprintf("%s:%s", ciphertext, iv)
	fmt.Println(result)
	fmt.Println("請保存上面的加密結果，解密時需要輸入完整字符串（包括冒號和IV部分）")
	fmt.Print("任意鍵繼續...")
	fmt.Scanln()
}

func Decrypt() {
	encryptedData := os.Getenv("STR")

	// 分離加密文本和IV
	parts := strings.Split(encryptedData, ":")
	if len(parts) != 2 {
		panic("加密數據格式錯誤，應為 'ciphertext:iv'")
	}

	ciphertext := parts[0]
	iv := parts[1]

	plaintext, err := aes.DecryptAES(ciphertext, iv, os.Getenv("AES_KEY"))
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(plaintext)
	fmt.Print("任意鍵繼續...")
	fmt.Scanln()
}
