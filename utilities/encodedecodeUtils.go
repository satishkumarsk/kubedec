package utilities

import (
	"crypto/rsa"
	"encoding/base64"
	"fmt"
)

type DecodedInput struct {
	Privatekey *rsa.PrivateKey
	Publickey  *rsa.PublicKey
	Data       map[string][]byte
	Label      []byte
}

func SecretDecoder(keyinput KeyInput, cipherinput CipherInput) (*DecodedInput, error) {
	var decodedresult DecodedInput

	decodedprvkey, _ := base64.StdEncoding.DecodeString(keyinput.PubPrivkey["tls.key"])
	decodedpubkey, _ := base64.StdEncoding.DecodeString(keyinput.PubPrivkey["tls.crt"])

	decodedresult.Privatekey, _ = PrivatekeyParser(decodedprvkey)
	decodedresult.Publickey, _ = PubkeyParser(decodedpubkey)
	decodedresult.Data = make(map[string][]byte)

	for k, v := range cipherinput.Spec.CipherData {

		data, _ := base64.StdEncoding.DecodeString(v)

		decodedresult.Data[k] = data
	}
	decodedresult.Label = []byte(fmt.Sprintf("%s/%s", cipherinput.Metadata["namespace"], cipherinput.Metadata["name"]))

	return &decodedresult, nil

}
