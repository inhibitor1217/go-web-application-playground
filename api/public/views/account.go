package views

import "github.com/inhibitor1217/go-web-application-playground/internal/features/account"

type AccountView struct {
	Id          string    `json:"id"`
	CreatedAt   Timestamp `json:"created_at"`
	UpdatedAt   Timestamp `json:"updated_at"`
	Email       string    `json:"email"`
	DisplayName string    `json:"display_name"`
	TouchedAt   Timestamp `json:"touched_at"`
}

func NewAccountView(a *account.Account) AccountView {
	return AccountView{
		Id:          a.Id,
		CreatedAt:   TimestampView(a.CreatedAt),
		UpdatedAt:   TimestampView(a.UpdatedAt),
		Email:       a.Email,
		DisplayName: a.DisplayName,
		TouchedAt:   TimestampView(a.TouchedAt),
	}
}
