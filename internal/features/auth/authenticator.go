package auth

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/env"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/jwt"
)

const (
	accessTokenCookie  = "x_access_token"
	refreshTokenCookie = "x_refresh_token"

	accessTokenTTL  = time.Duration(30) * time.Minute
	refreshTokenTTL = time.Duration(24) * time.Hour * 7
)

type Authenticator struct {
	env *env.Env
	jwt *jwt.Jwt
}

func NewAuthenticator(
	e *env.Env,
	j *jwt.Jwt,
) *Authenticator {
	return &Authenticator{
		env: e,
		jwt: j,
	}
}

func (a *Authenticator) SignAccessToken(p Principal) (string, error) {
	claims := a.jwt.DefaultClaimsBuilder().
		SetSubject(accessTokenSubject(p)).
		SetTTL(time.Now(), accessTokenTTL).
		Build()

	return a.jwt.Sign(claims)
}

func (a *Authenticator) SignRefreshToken(p Principal) (string, error) {
	claims := a.jwt.DefaultClaimsBuilder().
		SetSubject(refreshTokenSubject(p)).
		SetTTL(time.Now(), refreshTokenTTL).
		Build()

	return a.jwt.Sign(claims)
}

func (a *Authenticator) SignCookies(cx *gin.Context, p Principal) error {
	accessToken, err := a.SignAccessToken(p)
	if err != nil {
		return err
	}

	refreshToken, err := a.SignRefreshToken(p)
	if err != nil {
		return err
	}

	cx.SetCookie(accessTokenCookie, accessToken, int(accessTokenTTL.Seconds()), a.env.App.Domain, "", a.env.IsProduction(), true)
	cx.SetCookie(refreshTokenCookie, refreshToken, int(refreshTokenTTL.Seconds()), a.env.App.Domain, "", a.env.IsProduction(), true)

	return nil
}

func accessTokenSubject(p Principal) string {
	return fmt.Sprintf("auth:access_token:%s", p.Subject())
}

func isAccessTokenSubject(subject string) bool {
	return strings.HasPrefix(subject, "auth:access_token:")
}

func refreshTokenSubject(p Principal) string {
	return fmt.Sprintf("auth:refresh_token:%s", p.Subject())
}

func isRefreshTokenSubject(subject string) bool {
	return strings.HasPrefix(subject, "auth:refresh_token:")
}
