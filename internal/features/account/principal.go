package account

import "fmt"

type AccountPrincipal struct {
	Account Account
}

func NewPrincipal(a Account) *AccountPrincipal {
	return &AccountPrincipal{Account: a}
}

func (p *AccountPrincipal) Subject() string {
	return fmt.Sprintf("account:%s", p.Account.Id())
}
