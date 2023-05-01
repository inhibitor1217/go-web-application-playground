package account

import "fmt"

type AccountPrincipal struct {
	Account Account
}

func (p *AccountPrincipal) Subject() string {
	return fmt.Sprintf("account:%s", p.Account.Id())
}
