package cryptopus_test

import (
	"crypto/x509"
	"cryptopus"
	"os"
	"reflect"
	"testing"
)

func TestKeyGen(t *testing.T) {
	const TEST_DIR = "D:\\GOPHER\\goframe\\goframe\\inter\\cryptopus"
	file, err := os.CreateTemp(TEST_DIR, "*.pem")
	if err != nil {
		t.Error("Error with temp fil creating: ", err)
	}
	defer file.Close()
	key, err := cryptopus.GenerateRSAPrivateKey()
	if err != nil {
		t.Error("Fail to generate RSAP key", err)
	}
	if err = cryptopus.WriteKeyToFile(key, file); err != nil {
		t.Error("Error with writing key to file", err)
	}
	file.Close()

	fileName := file.Name()
	newFile, err := os.Open(fileName)

	if err != nil {
		t.Error("Error with file open ", err)
	}
	defer newFile.Close()

	keyFromFile, err := cryptopus.ReadKeyFromFile(newFile)
	if err != nil {
		t.Error("Error with reading key from file: ", err)
	}
	k1 := x509.MarshalPKCS1PrivateKey(keyFromFile)
	if k1 == nil {
		t.Error("k1 was empty")
	}
	k2 := x509.MarshalPKCS1PrivateKey(key)
	if k2 == nil {
		t.Error("k2 was empty")
	}
	if !reflect.DeepEqual(k1, k2) {
		t.Log("k1 ", k1)
		t.Log("k2 ", k2)
		t.Error("Fail! k1 and k2 not equal")
	}
	t.Log("Passed")
}
