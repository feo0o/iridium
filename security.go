package iridium

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
)

func parseKey(key string) (keyBytes [32]byte, err error) {
	if len(key) < 6 {
		return keyBytes, errors.New("key too short")
	}
	return sha256.Sum256([]byte(key)), nil
}

// AESEncryptWithCFB encrypt string with AES CFB mode, output hex string
func AESEncryptWithCFB(plain, key string) (secret string, err error) {
	keyBytes, err := parseKey(key)
	if err != nil {
		return secret, err
	}
	plainBytes := []byte(plain)

	block, err := aes.NewCipher(keyBytes[:])
	if err != nil {
		return secret, err
	}
	secretBytes := make([]byte, aes.BlockSize+len(plainBytes))
	iv := secretBytes[:aes.BlockSize]
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return secret, err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(secretBytes[aes.BlockSize:], plainBytes)
	return hex.EncodeToString(secretBytes), nil
}

// AESDecryptWithCFB decrypt hex string with AES CFB mode
func AESDecryptWithCFB(secret, key string) (plain string, err error) {
	keyBytes, err := parseKey(key)
	if err != nil {
		return plain, err
	}
	secretBytes, err := hex.DecodeString(secret)
	if err != nil {
		return plain, err
	}
	if len(secretBytes) < aes.BlockSize {
		return plain, errors.New("secret too short")
	}

	block, err := aes.NewCipher(keyBytes[:])
	if err != nil {
		return plain, err
	}
	iv := secretBytes[:aes.BlockSize]
	secretBytes = secretBytes[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(secretBytes, secretBytes)
	return string(secretBytes), nil
}
