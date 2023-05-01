package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inhibitor1217/go-web-application-playground/api/public/lib"
	"github.com/inhibitor1217/go-web-application-playground/api/public/views"
	"github.com/inhibitor1217/go-web-application-playground/internal/features/account"
	"github.com/inhibitor1217/go-web-application-playground/internal/features/auth"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/log"
)

// SignUp godoc
//
//	@Summary		Sign up (register)
//	@Description	Registers a new account.
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			body	body		auth.SignUp.request	true	"Request body"
//	@Success		201		{object}	auth.SignUp.ok
//	@Failure		400		{object}	views.ErrorView[auth.SignUp.accountExists]
//	@Router			/auth/sign-up [post]
func (h *Handler) SignUp(cx *gin.Context) {
	type request struct {
		Email       string  `json:"email" binding:"required,email"`
		Password    string  `json:"password" binding:"required,min=8,max=32"`
		DisplayName *string `json:"display_name" binding:"omitempty,min=1,max=256"`
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

	exists, err := h.accountSvc.ExistsOfEmail(cx.Request.Context(), req.Email)
	if err != nil {
		lib.Panic(cx, err, h.l)
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

	a, err := h.accountSvc.Create(cx.Request.Context(), &account.CreateDTO{
		Email:       req.Email,
		Password:    req.Password,
		DisplayName: req.DisplayName,
	})

	if err != nil {
		lib.Panic(cx, err, h.l)
		return
	}

	cx.JSON(http.StatusCreated, ok{
		Account: views.NewAccountView(a),
	})
}

// SignIn godoc
//
//	@Summary		Sign in (login)
//	@Description	Signs in to an account using email and password.
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			body	body		auth.SignIn.request	true	"Request body"
//	@Success		200		{object}	auth.SignUp.ok
//	@Failure		401
//	@Router			/auth/sign-in [post]
func (h *Handler) SignIn(cx *gin.Context) {
	type request struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	type ok struct {
		Account views.AccountView `json:"account"`
	}

	req := request{}
	if err := cx.ShouldBindJSON(&req); err != nil {
		views.ValidationError(cx, err)
		return
	}

	a, err := h.accountSvc.FindByEmail(cx.Request.Context(), req.Email)
	if err != nil {
		lib.Panic(cx, err, h.l)
		return
	}
	if a == nil {
		cx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	pass, err := account.ValidatePassword(a, req.Password)

	if err != nil {
		lib.Panic(cx, err, h.l)
		return
	}
	if !pass {
		cx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	principal := account.NewPrincipal(a)
	if err := h.auth.Sign(cx, principal); err != nil {
		lib.Panic(cx, err, h.l)
		return
	}

	cx.JSON(http.StatusOK, ok{
		Account: views.NewAccountView(a),
	})

	h.l.Debug("Signed in", log.String("account.id", a.Id()))
}

// Touch godoc
//
//	@Summary		Touch
//	@Description	Touches the session and renews tokens.
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	auth.Touch.ok
//	@Router			/auth/touch [post]
func (h *Handler) Touch(cx *gin.Context) {
	type ok struct {
		Principal views.PrincipalView `json:"principal"`
	}

	principal, exists := cx.Get(lib.Principal)
	if !exists {
		cx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	cx.JSON(http.StatusOK, ok{
		Principal: views.NewPrincipalView(principal.(auth.Principal)),
	})
}

// SignOut godoc
//
//	@Summary		Sign out (logout)
//	@Description	Signs out from the account session.
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Success		204
//	@Router			/auth/sign-out [post]
func (h *Handler) SignOut(cx *gin.Context) {
	cx.JSON(http.StatusNotImplemented, "Not implemented")
}
