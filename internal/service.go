package internal

import (
	"fmt"
	"main/internal/aes"
	"main/internal/config"
	"os"
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
	if str, err := aes.EncryptAES(os.Getenv("STR"), os.Getenv("AES_IV"), os.Getenv("AES_KEY")); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(str)
		fmt.Print("任意鍵繼續...")
		fmt.Scanln()
	}
}

func Decrypt() {
	if str, err := aes.DecryptAES(os.Getenv("STR"), os.Getenv("AES_IV"), os.Getenv("AES_KEY")); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(str)
		fmt.Print("任意鍵繼續...")
		fmt.Scanln()
	}
}
