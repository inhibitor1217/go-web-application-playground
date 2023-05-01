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
	"github.com/pkg/errors"
)

const (
	accessTokenCookie  = "x_access_token"
	refreshTokenCookie = "x_refresh_token"

	accessTokenTTL  = time.Duration(30) * time.Minute
	refreshTokenTTL = time.Duration(24) * time.Hour * 7
)

var AccessTokenNotFound = errors.New("access token not found")
var AccessTokenInvalid = errors.New("invalid access token")
var InvalidPrincipal = errors.New("invalid principal")

type Authenticator struct {
	accountSvc account.Service
	env        *env.Env
	jwt        *jwt.Jwt
}

func NewAuthenticator(
	accountSvc account.Service,
	e *env.Env,
	j *jwt.Jwt,
) *Authenticator {
	return &Authenticator{
		accountSvc: accountSvc,
		env:        e,
		jwt:        j,
	}
}

func (a *Authenticator) SignCookies(cx *gin.Context, p Principal) error {
	accessToken, err := a.signAccessToken(p)
	if err != nil {
		return err
	}

	refreshToken, err := a.signRefreshToken(p)
	if err != nil {
		return err
	}

	cx.SetCookie(accessTokenCookie, accessToken, int(accessTokenTTL.Seconds()), "/", a.env.App.Domain, a.env.IsProduction(), true)
	cx.Header("X-Refresh-Token", refreshToken)

	return nil
}

func (a *Authenticator) signAccessToken(p Principal) (string, error) {
	claims := a.jwt.DefaultClaimsBuilder().
		SetSubject(a.accessTokenSubject(p)).
		SetTTL(time.Now(), accessTokenTTL).
		Build()

	return a.jwt.Sign(claims)
}

func (a *Authenticator) signRefreshToken(p Principal) (string, error) {
	claims := a.jwt.DefaultClaimsBuilder().
		SetSubject(a.refreshTokenSubject(p)).
		SetTTL(time.Now(), refreshTokenTTL).
		Build()

	return a.jwt.Sign(claims)
}

func (a *Authenticator) Authenticate(cx *gin.Context) (Principal, error) {
	accessToken, err := cx.Cookie(accessTokenCookie)
	if err == http.ErrNoCookie {
		cx.AbortWithStatus(http.StatusUnauthorized)
		return nil, AccessTokenNotFound
	} else if err != nil {
		return nil, err
	}

	return a.authenticateFromAccessToken(cx.Request.Context(), accessToken)
}

func (a *Authenticator) authenticateFromAccessToken(cx context.Context, accessToken string) (Principal, error) {
	claims, err := a.jwt.Parse(accessToken)
	if err != nil {
		return nil, AccessTokenInvalid
	}

	if !a.isValidAccessTokenClaims(claims) {
		return nil, AccessTokenInvalid
	}

	return a.makePrincipal(cx, claims.Subject)
}

func (a *Authenticator) makePrincipal(cx context.Context, subject string) (Principal, error) {
	principalSubject := strings.TrimPrefix(subject, "auth:access_token:")
	paths := strings.Split(principalSubject, ":")

	if len(paths) != 2 {
		return nil, AccessTokenInvalid
	}

	switch paths[0] {
	case "account":
		a, err := a.accountSvc.Find(cx, paths[1])
		if err != nil {
			return nil, err
		}
		if a == nil {
			return nil, InvalidPrincipal
		}
		return account.NewPrincipal(a), nil
	default:
		return nil, InvalidPrincipal
	}
}

func (a *Authenticator) accessTokenSubject(p Principal) string {
	return fmt.Sprintf("auth:access_token:%s", p.Subject())
}

func (a *Authenticator) refreshTokenSubject(p Principal) string {
	return fmt.Sprintf("auth:refresh_token:%s", p.Subject())
}

func (a *Authenticator) isValidAccessTokenClaims(claims *jwt.Claims) bool {
	if !strings.HasPrefix(claims.Subject, "auth:access_token:") {
		return false
	}
	if claims.Issuer != a.jwt.Issuer() {
		return false
	}
	return true
}

func (a *Authenticator) isValidRefreshTokenClaims(claims *jwt.Claims) bool {
	if !strings.HasPrefix(claims.Subject, "auth:refresh_token:") {
		return false
	}
	if claims.Issuer != a.jwt.Issuer() {
		return false
	}
	return true
}
