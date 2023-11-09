package cryptopus

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"path/filepath"
)

func GenerateRSAPrivateKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}

func CryptopusMain() {
	args := os.Args
	arguments := len(args)

	if arguments < 2 {
		fmt.Println("Not enought argements. Maybe you forgot path to output?")
		os.Exit(1)
	}

	file := args[1]
	fileExtension := filepath.Ext(file)
	// TODO: расширение файла .pem а не pem
	if fileExtension != "pem" {
		fmt.Println("You are made of stupid. file extension mist be a pem. Your extansion is ", fileExtension)
	}

	key, err := GenerateRSAPrivateKey()
	if err != nil {
		fmt.Println("Erorr with key generating: ", err)
		os.Exit(1)
	}
	privateKeyFile, err := os.Create(file)
	if err != nil {
		fmt.Println("Some shit hapen. I don't know what kind of shit, so you can find out yourself: ", err)
		os.Exit(1)
	}
	defer privateKeyFile.Close()

	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	err = pem.Encode(privateKeyFile, privateKeyPEM)
	if err != nil {
		fmt.Println("Can't write key in file: ", err)
		os.Exit(1)
	}

	fmt.Println("Suckes, key created and writed at ", file)
}
