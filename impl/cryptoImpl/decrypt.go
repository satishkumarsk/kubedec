package cryptoImpl

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"io"
)

const (
	sessionKeyBytes = 32
)

var ErrTooShort = errors.New("SealedSecret data is too short")

func Decrypter(rnd io.Reader, privKey *rsa.PrivateKey, ciphertext, label []byte) ([]byte, error) {

	if len(ciphertext) < 2 {
		return nil, ErrTooShort
	}
	rsaLen := int(binary.BigEndian.Uint16(ciphertext))
	if len(ciphertext) < rsaLen+2 {
		return nil, ErrTooShort
	}

	rsaCiphertext := ciphertext[2 : rsaLen+2]
	aesCiphertext := ciphertext[rsaLen+2:]

	sessionKey, err := rsa.DecryptOAEP(sha256.New(), rnd, privKey, rsaCiphertext, label)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(sessionKey)
	if err != nil {
		return nil, err
	}

	aed, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Key is only used once, so zero nonce is ok
	zeroNonce := make([]byte, aed.NonceSize())

	plaintext, err := aed.Open(nil, zeroNonce, aesCiphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
