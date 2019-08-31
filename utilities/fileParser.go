package utilities

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

type EncryptedData struct {
	CipherData map[string]string `yaml:"encryptedData"`
}
type CipherInput struct {
	Metadata map[string]string `yaml:"metadata"`
	Spec     EncryptedData     `yaml:"spec"`
}

type KeyInput struct {
	PubPrivkey map[string]string `yaml:"data"`
}

func ParseCipherYaml(sealedyamlpath string) *CipherInput {

	var cinput CipherInput

	sealedyamlpath, _ = filepath.Abs(sealedyamlpath)

	ciphercontent, err := ioutil.ReadFile(sealedyamlpath)
	if err != nil {
		fmt.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(ciphercontent, &cinput)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return &cinput

}

func ParsemasterkeyYaml(masterkeypath string) *KeyInput {

	var keyinput KeyInput

	masterkeypath, _ = filepath.Abs(masterkeypath)

	masterkeycontent, err := ioutil.ReadFile(masterkeypath)
	if err != nil {
		fmt.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(masterkeycontent, &keyinput)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return &keyinput

}
