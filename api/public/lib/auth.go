package lib

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inhibitor1217/go-web-application-playground/internal/features/auth"
)

func RequireAuth(cx *gin.Context) (auth.Principal, bool) {
	principal, exists := cx.Get(Principal)
	if !exists {
		cx.AbortWithStatus(http.StatusUnauthorized)
		return nil, false
	}
	return principal.(auth.Principal), true
}
