package auth

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
)

type Body struct {
	IAT int64
	ISS string
	SLV string
	UID string
	RO  string
}

func ImportKey(path string) (*rsa.PrivateKey, error) {
	s, err := ioutil.ReadFile(path)
	if err != nil {
		panic("Key Not Available")
	}
	block, _ := pem.Decode(s)
	pvk, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return pvk, nil
}
