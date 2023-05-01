package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/inhibitor1217/go-web-application-playground/internal/features/account"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/env"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/jwt"
)

const (
	accessTokenCookie  = "x_access_token"
	refreshTokenHeader = "X-Refresh-Token"

	accessTokenTTL  = time.Duration(30) * time.Minute
	refreshTokenTTL = time.Duration(24) * time.Hour * 7
)

type tokenAuthenticator struct {
	accountSvc account.Service
	env        *env.Env
	jwt        *jwt.Jwt
}

func (a *tokenAuthenticator) Sign(cx *gin.Context, p Principal) error {
	accessToken, err := a.signAccessToken(p)
	if err != nil {
		return err
	}

	refreshToken, err := a.signRefreshToken(p)
	if err != nil {
		return err
	}

	cx.SetCookie(accessTokenCookie, accessToken, int(accessTokenTTL.Seconds()), "/", a.env.App.Domain, a.env.IsProduction(), true)
	cx.Header(refreshTokenHeader, refreshToken)

	return nil
}

func (a *tokenAuthenticator) Authenticate(cx *gin.Context) (Principal, error) {
	accessToken, err := cx.Cookie(accessTokenCookie)
	if err == http.ErrNoCookie {
		return nil, AuthRequired
	} else if err != nil {
		return nil, err
	}

	return a.authenticateFromAccessToken(cx.Request.Context(), accessToken)
}

func (a *tokenAuthenticator) WillExpire(cx *gin.Context) bool {
	accessToken, err := cx.Cookie(accessTokenCookie)
	if err == http.ErrNoCookie {
		return false
	} else if err != nil {
		return false
	}

	claims, err := a.jwt.Parse(accessToken)
	if err != nil {
		return false
	}

	if claims.WillExpireIn(accessTokenTTL / 2) {
		return true
	}

	return false
}

func (a *tokenAuthenticator) Refresh(cx *gin.Context) (Principal, error) {
	refreshToken := cx.GetHeader(refreshTokenHeader)
	if refreshToken == "" {
		return nil, AuthRequired
	}

	principal, err := a.authenticateFromRefreshToken(cx.Request.Context(), refreshToken)
	if err != nil {
		return nil, err
	}

	accessToken, err := a.signAccessToken(principal)
	if err != nil {
		return nil, err
	}
	refreshToken, err = a.signRefreshToken(principal)
	if err != nil {
		return nil, err
	}

	cx.SetCookie(accessTokenCookie, accessToken, int(accessTokenTTL.Seconds()), "/", a.env.App.Domain, a.env.IsProduction(), true)
	cx.Header(refreshTokenHeader, refreshToken)

	return principal, nil
}

func (a *tokenAuthenticator) Clear(cx *gin.Context) error {
	cx.SetCookie(accessTokenCookie, "", 0, "/", a.env.App.Domain, a.env.IsProduction(), true)

	return nil
}

func (a *tokenAuthenticator) signAccessToken(p Principal) (string, error) {
	claims := a.jwt.DefaultClaimsBuilder().
		SetSubject(a.accessTokenSubject(p)).
		SetTTL(time.Now(), accessTokenTTL).
		Build()

	return a.jwt.Sign(claims)
}

func (a *tokenAuthenticator) signRefreshToken(p Principal) (string, error) {
	claims := a.jwt.DefaultClaimsBuilder().
		SetSubject(a.refreshTokenSubject(p)).
		SetTTL(time.Now(), refreshTokenTTL).
		Build()

	return a.jwt.Sign(claims)
}

func (a *tokenAuthenticator) authenticateFromAccessToken(cx context.Context, accessToken string) (Principal, error) {
	claims, err := a.jwt.Parse(accessToken)
	if err != nil {
		return nil, AuthRequired
	}

	if !a.isValidAccessTokenClaims(claims) {
		return nil, AuthRequired
	}

	return a.makePrincipal(cx, claims.Subject)
}

func (a *tokenAuthenticator) authenticateFromRefreshToken(cx context.Context, refreshToken string) (Principal, error) {
	claims, err := a.jwt.Parse(refreshToken)
	if err != nil {
		return nil, AuthRequired
	}

	if !a.isValidRefreshTokenClaims(claims) {
		return nil, AuthRequired
	}

	return a.makePrincipal(cx, claims.Subject)
}

func (a *tokenAuthenticator) makePrincipal(cx context.Context, subject string) (Principal, error) {
	principalSubject := strings.TrimPrefix(subject, "auth:access_token:")
	paths := strings.Split(principalSubject, ":")

	if len(paths) != 2 {
		return nil, AuthRequired
	}

	switch paths[0] {
	case "account":
		a, err := a.accountSvc.Find(cx, paths[1])
		if err != nil {
			return nil, err
		}
		if a == nil {
			return nil, AuthRequired
		}
		return account.NewPrincipal(a), nil
	default:
		return nil, AuthRequired
	}
}

func (a *tokenAuthenticator) accessTokenSubject(p Principal) string {
	return fmt.Sprintf("auth:access_token:%s:%s", p.Type(), p.Id())
}

func (a *tokenAuthenticator) refreshTokenSubject(p Principal) string {
	return fmt.Sprintf("auth:refresh_token:%s:%s", p.Type(), p.Id())
}

func (a *tokenAuthenticator) isValidAccessTokenClaims(claims *jwt.Claims) bool {
	if !strings.HasPrefix(claims.Subject, "auth:access_token:") {
		return false
	}
	if claims.Issuer != a.jwt.Issuer() {
		return false
	}
	return true
}

func (a *tokenAuthenticator) isValidRefreshTokenClaims(claims *jwt.Claims) bool {
	if !strings.HasPrefix(claims.Subject, "auth:refresh_token:") {
		return false
	}
	if claims.Issuer != a.jwt.Issuer() {
		return false
	}
	return true
}
