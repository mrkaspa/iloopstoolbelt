package command

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

var (
	gcm   cipher.AEAD
	nonce []byte
)

func init() {
	key := []byte("yqw134ms5b5ar1Me") // any 128-, 192-, or 256-bit key
	b, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	gcm, err = cipher.NewGCM(b)
	if err != nil {
		panic(err)
	}
	nonce = make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		panic(err)
	}
}

func cypher(data []byte) ([]byte, error) {
	ciphertext := gcm.Seal(nil, nonce, data, nil)
	return ciphertext, nil
}

func decypher(data []byte) ([]byte, error) {
	text, err := gcm.Open(nil, nonce, data, nil)
	if err != nil {
		return nil, err
	}
	return text, nil
}
