package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/inhibitor1217/go-web-application-playground/api/public/lib"
	"github.com/inhibitor1217/go-web-application-playground/internal/features/auth"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/log"
)

type Auth struct {
	auth auth.Authenticator
	l    *log.Logger
}

func NewAuth(
	a auth.Authenticator,
	l *log.Logger,
) *Auth {
	return &Auth{
		auth: a,
		l:    l,
	}
}

func (a *Auth) Handler() gin.HandlerFunc {
	return func(cx *gin.Context) {
		principal, err := a.auth.Authenticate(cx)

		if err == auth.AuthRequired {
			// Access token is invalid, but we can continue with refresh token
			principal, err = a.auth.Refresh(cx)

			if err == auth.AuthRequired {
				// Refresh token is invalid, so we need to sign in again
				// Continue with no principal
				// TODO Implement anonymous principal using session cookies
				cx.Next()
				return
			} else if err != nil {
				lib.Panic(cx, err, a.l)
				return
			} else {
				// Refreshed
				if principal != nil {
					cx.Set(lib.Principal, principal)
				}
				cx.Next()
				return
			}
		} else if err != nil {
			lib.Panic(cx, err, a.l)
			return
		} else {
			// Authenticated
			if principal != nil {
				cx.Set(lib.Principal, principal)
			}
			cx.Next()
			return
		}
	}
}
