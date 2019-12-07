package cryptoImpl

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"
)

func Test_Decrypter(t *testing.T) {
	// Generate random pub/private key combination
	key, _ := rsa.GenerateKey(rand.Reader, 1024)
	randomSeedText := "randomtext"
	tests := []struct {
		name      string
		plaintext string
		want      string
	}{
		{
			"Decrypt foo",
			"foo",
			"foo",
		},
		{
			"Decrypt bar",
			"helloworld",
			"helloworld",
		},
		{
			"Decrypt int",
			"123456",
			"123456",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			pubkey := key.Public().(*rsa.PublicKey)
			cipherText, err := HybridEncrypt(rand.Reader, pubkey, []byte(test.plaintext), []byte(randomSeedText))
			if err != nil {
				t.Fatalf("Encryption failed with error %s", err.Error())

			}
			decryptedtext, err := Decrypter(rand.Reader, key, cipherText, []byte(randomSeedText))
			if err != nil {
				t.Fatalf("Decryption failed for test %s", err.Error())

			}
			if string(decryptedtext) != test.want {
				t.Errorf(" Test %s want text as %s got %s", test.name, test.want, string(decryptedtext))

			}

		})

	}

}
