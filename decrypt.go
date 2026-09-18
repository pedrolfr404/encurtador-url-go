package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func decrypt(encryptedUrl string) string {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	secretKey := os.Getenv("secretKey")

	block, err := aes.NewCipher([]byte(secretKey))

	if err != nil {
		log.Fatal(err)
	}

	cipherText, err := hex.DecodeString(encryptedUrl)
	if err != nil {
		log.Fatal(err)
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText)
}
