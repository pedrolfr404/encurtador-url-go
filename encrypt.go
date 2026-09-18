package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func encrypt(originalUrl string) string {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	secretKey := os.Getenv("secretKey")

	block, err := aes.NewCipher([]byte(secretKey))

	if err != nil {
		log.Fatal(err)
	}

	plainText := []byte(originalUrl)
	cipherText := make([]byte, aes.BlockSize+len(plainText))

	iv := cipherText[:aes.BlockSize]

	if _, err := rand.Read(iv); err != nil {
		log.Fatal(err)
	}

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	return hex.EncodeToString(cipherText)
}
