package secret

import (
	"acme-manager/logger"
	"acme-manager/util"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"golang.org/x/crypto/chacha20poly1305"
)

type XChaCha20 struct {
	aead cipher.AEAD
}

func (x *XChaCha20) Encrypt(data string) (string, error) {
	nonce := make([]byte, chacha20poly1305.NonceSizeX)
	if _, err := rand.Read(nonce); err != nil {
		return "", err
	}

	encrypted := x.aead.Seal(nonce, nonce, []byte(data), nil)
	encoded := base64.StdEncoding.EncodeToString(encrypted)
	return encoded, nil
}

func (x *XChaCha20) Decrypt(data string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	nonceSize := chacha20poly1305.NonceSizeX
	if len(data) < nonceSize {
		return "", nil
	}

	nonce, cipherText := decoded[:nonceSize], decoded[nonceSize:]
	decrypted, err := x.aead.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(decrypted), nil
}

func NewXChaCha2Provider(key string) *XChaCha20 {
	keyBytes, err := hex.DecodeString(key)
	if err != nil {
		logger.Fatalf("Failed to decode XChaCha20 key %s: %s", util.MakeSensitive(key), err)
	}
	aead, err := chacha20poly1305.NewX(keyBytes)
	if err != nil {
		logger.Fatalf("Failed to create XChaCha20 provider: %s", err)
	}
	return &XChaCha20{aead: aead}
}
