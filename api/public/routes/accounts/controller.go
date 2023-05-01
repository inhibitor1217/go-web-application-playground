package accounts

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inhibitor1217/go-web-application-playground/api/public/lib"
	"github.com/inhibitor1217/go-web-application-playground/api/public/views"
	"github.com/inhibitor1217/go-web-application-playground/internal/features/account"
)

// Me godoc
//
//	@Summary		Get my account
//	@Description	Gets the account of the authenticated user.
//	@Tags			Accounts
//	@Accept			json
//	@Produce		json
//	@Security		AccountPrincipal
//	@Success		200	{object}	accounts.Me.ok
//	@Router			/accounts/me [get]
func (h *Handler) Me(cx *gin.Context) {
	type ok struct {
		Account views.AccountView `json:"account"`
	}

	principal, pass := lib.RequireAuth(cx)
	if !pass {
		return
	}

	accountPrincipal, pass := principal.(*account.AccountPrincipal)
	if !pass {
		cx.AbortWithStatus(http.StatusForbidden)
		return
	}

	cx.JSON(http.StatusOK, ok{
		Account: views.NewAccountView(accountPrincipal.Account),
	})
}
