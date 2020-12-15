package auth

import (
	"crypto/rsa"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
)

func GenerateJWT(privateKey *rsa.PrivateKey, b Body) ([]byte, error) {
	t := jwt.New()

	err := t.Set(jwt.IssuerKey, b.ISS)
	err = t.Set(jwt.IssuedAtKey, b.IAT)
	err = t.Set(`slv`, b.SLV)
	err = t.Set(`uid`, b.UID)
	err = t.Set(`ro`, b.RO)

	signed, err := jwt.Sign(t, jwa.RS256, privateKey)
	if err != nil {
		return nil, err
	}

	return signed, nil
}
