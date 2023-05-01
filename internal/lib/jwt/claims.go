package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
}

type ClaimsBuilder struct {
	claims *Claims
}

func NewClaimsBuilder() *ClaimsBuilder {
	return &ClaimsBuilder{
		claims: &Claims{},
	}
}

func (b *ClaimsBuilder) Build() *Claims {
	return b.claims
}

func (b *ClaimsBuilder) SetIssuer(issuer string) *ClaimsBuilder {
	b.claims.Issuer = issuer
	return b
}

func (b *ClaimsBuilder) SetSubject(subject string) *ClaimsBuilder {
	b.claims.Subject = subject
	return b
}

func (b *ClaimsBuilder) SetExpiresAt(expiresAt time.Time) *ClaimsBuilder {
	b.claims.ExpiresAt = jwt.NewNumericDate(expiresAt)
	return b
}

func (b *ClaimsBuilder) SetIssuedAt(issuedAt time.Time) *ClaimsBuilder {
	b.claims.IssuedAt = jwt.NewNumericDate(issuedAt)
	return b
}

func (b *ClaimsBuilder) SetTTL(issuedAt time.Time, ttl time.Duration) *ClaimsBuilder {
	b.claims.IssuedAt = jwt.NewNumericDate(issuedAt)
	b.claims.ExpiresAt = jwt.NewNumericDate(issuedAt.Add(ttl))
	return b
}
