package util

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"github.com/cockroachdb/errors"
)

func PrivateKeyToPkcs8Pem(privateKey crypto.PrivateKey) (string, error) {
	pkcs8PrivateKeyBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return "", errors.WithStack(err)
	}
	privateKeyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: pkcs8PrivateKeyBytes,
	})
	return string(privateKeyPem), nil
}

func Pkcs8PemToPrivateKey(pemString string) (crypto.Signer, error) {
	block, _ := pem.Decode([]byte(pemString))
	if block == nil {
		return nil, errors.New("failed to decode PEM block")
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	signer, ok := privateKey.(crypto.Signer)
	if !ok {
		return nil, errors.New("private key is not a signer")
	}
	return signer, nil
}

func PublicKeyFromPrivateKey(privateKey crypto.PrivateKey) (crypto.PublicKey, error) {
	switch k := privateKey.(type) {
	case *rsa.PrivateKey:
		return &k.PublicKey, nil
	case *ecdsa.PrivateKey:
		return &k.PublicKey, nil
	default:
		return nil, errors.New("unsupported private key type")
	}
}

func PublicKeyFingerprint(publicKey crypto.PublicKey) (string, error) {
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", errors.WithStack(err)
	}
	hash := sha256.Sum256(publicKeyBytes)
	return hex.EncodeToString(hash[:]), nil
}

func PrivateKeyFingerprint(privateKey crypto.PrivateKey) (string, error) {
	publicKey, err := PublicKeyFromPrivateKey(privateKey)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return PublicKeyFingerprint(publicKey)
}

func ParseX509Certificate(certPem string) (*x509.Certificate, error) {
	block, _ := pem.Decode([]byte(certPem))
	if block == nil {
		return nil, errors.New("failed to decode PEM block")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return cert, nil
}
