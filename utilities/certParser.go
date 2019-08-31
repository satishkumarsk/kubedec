package utilities

import (
	"crypto/rsa"
	"fmt"

	certUtil "k8s.io/client-go/util/cert"
)

func PrivatekeyParser(prvkey []byte) (*rsa.PrivateKey, error) {

	key, err := certUtil.ParsePrivateKeyPEM(prvkey)

	if err != nil {

		return nil, err
	}
	return key.(*rsa.PrivateKey), nil
}

func PubkeyParser(pubkey []byte) (*rsa.PublicKey, error) {

	certs, err := certUtil.ParseCertsPEM(pubkey)

	if err != nil {

		return nil, err
	}
	cert, ok := certs[0].PublicKey.(*rsa.PublicKey)

	if !ok {

		return nil, fmt.Errorf("data doesn't contain valid RSA Public Key")
	}
	return cert, nil

}
