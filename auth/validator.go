package auth

import (
	"crypto/rsa"
	"encoding/json"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
)

func ValidateJWT(tokenString []byte, privateKey *rsa.PrivateKey) (Body, error) {
	var b Body
	token, err := jwt.Parse(tokenString, jwt.WithVerify(jwa.RS256, &privateKey.PublicKey))
	if err != nil {
		return Body{}, err
	}
	buf, err := json.MarshalIndent(token, "", "  ")
	if err != nil {
		return Body{}, err
	}
	err = json.Unmarshal(buf, &b)
	if err != nil {
		return Body{}, err
	}
	return b, nil
}
