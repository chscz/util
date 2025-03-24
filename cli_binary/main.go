package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

const key = "this-is-24-byte-test-key"

func main() {
	f := flag.NewFlagSet("Test CLI Options", flag.ExitOnError)
	var mode int
	var message string

	f.IntVar(&mode, "mode", -1, "required: [0: encrypt, 1: decrypt]]")
	f.StringVar(&message, "message", "", "required: string")
	if err := f.Parse(os.Args[1:]); err != nil {
		panic(err)
	}

	if mode == -1 || message == "" {
		fmt.Println("must provide either [encrypt] or [decrypt]")
		fmt.Println(`ex> ./aes_cli -mode 0 -message "hi"`)
		return
	}

	if message == "" {
		fmt.Println("must provide [message]")
		fmt.Println(`ex> ./aes_cli -mode 0 -message "hi"`)
		return
	}

	var result string
	var err error
	switch mode {
	case 0:
		result, err = encrypt([]byte(message))
	case 1:
		result, err = decrypt(message)
	default:
		result, err = "", errors.New("invalid mode")
	}

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

func encrypt(text []byte) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := aesgcm.Seal(nonce, nonce, text, nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func decrypt(text string) (string, error) {
	decodedCiphertext, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesgcm.NonceSize()
	nonce, ciphertext := decodedCiphertext[:nonceSize], decodedCiphertext[nonceSize:]

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
