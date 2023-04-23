package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inhibitor1217/go-web-application-playground/api/public/views"
	"github.com/inhibitor1217/go-web-application-playground/internal/features/account"
)

// SignUp godoc
//
//	@Summary		Sign up
//	@Description	Register a new account
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			body	body		auth.SignUp.request	true	"Request body"
//	@Success		201		{object}	auth.SignUp.ok
//	@Failure		400		{object}	views.ErrorView[auth.SignUp.accountExists]
//	@Router			/auth/sign-up [post]
func (h *Handler) SignUp(cx *gin.Context) {
	type request struct {
		Email       string `json:"email" binding:"required,email"`
		Password    string `json:"password" binding:"required,min=8,max=32"`
		DisplayName string `json:"display_name" binding:"min=1,max=256"`
	}

	type ok struct {
		Account views.AccountView `json:"account"`
	}

	type accountExists struct {
		Email string `json:"email"`
	}

	req := request{}
	if err := cx.ShouldBindJSON(&req); err != nil {
		views.ValidationError(cx, err)
		return
	}

	exists, err := h.accountSvc.ExistsOfEmail(req.Email)
	if err != nil {
		views.Panic(cx, err)
		return
	}

	if exists {
		views.ClientError(cx, views.ErrorView[accountExists]{
			Type:    "account_exists",
			Message: "Account already exists",
			Payload: accountExists{
				Email: req.Email,
			},
		})
		return
	}

	a, err := h.accountSvc.Create(&account.CreateDTO{
		Email:       req.Email,
		Password:    req.Password,
		DisplayName: req.DisplayName,
	})

	if err != nil {
		views.Panic(cx, err)
		return
	}

	cx.JSON(http.StatusCreated, ok{
		Account: views.NewAccountView(a),
	})
}

func (h *Handler) SignIn(cx *gin.Context) {
	cx.JSON(http.StatusNotImplemented, "Not implemented")
}

func (h *Handler) Touch(cx *gin.Context) {
	cx.JSON(http.StatusNotImplemented, "Not implemented")
}

func (h *Handler) SignOut(cx *gin.Context) {
	cx.JSON(http.StatusNotImplemented, "Not implemented")
}
