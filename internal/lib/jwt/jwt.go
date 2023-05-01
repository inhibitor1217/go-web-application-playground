package jwt

import (
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/env"
	"github.com/pkg/errors"
)

type Jwt struct {
	issuer string
	secret []byte
}

func NewJwt(e *env.Env) (*Jwt, error) {
	secretStr := e.Auth.JwtSecret
	if secretStr == "" {
		return nil, errors.New("JWT secret is not set")
	}

	return &Jwt{
		issuer: e.App.Name,
		secret: []byte(secretStr),
	}, nil
}

func (j *Jwt) Sign(claims *Claims) (string, error) {
	return Sign(j.secret, claims)
}

func (j *Jwt) Parse(tokenString string) (*Claims, error) {
	return Parse(j.secret, tokenString)
}

func (j *Jwt) DefaultClaimsBuilder() *ClaimsBuilder {
	return NewClaimsBuilder().
		SetIssuer(j.issuer)
}
