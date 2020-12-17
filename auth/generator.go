package auth

import (
	"crypto/rsa"
	"fmt"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
)

func GenerateJWT(privateKey *rsa.PrivateKey, b Body) (string, error) {
	t := jwt.New()

	err := t.Set(jwt.IssuerKey, b.ISS)
	err = t.Set(jwt.IssuedAtKey, b.IAT)
	err = t.Set(`uid`, b.UID)
	err = t.Set(`ro`, b.RO)

	signed, err := jwt.Sign(t, jwa.RS256, privateKey)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", signed), nil
}
