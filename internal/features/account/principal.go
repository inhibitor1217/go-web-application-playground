package account

type AccountPrincipal struct {
	Account Account
}

func NewPrincipal(a Account) *AccountPrincipal {
	return &AccountPrincipal{Account: a}
}

func (p *AccountPrincipal) Type() string {
	return "account"
}

func (p *AccountPrincipal) Id() string {
	return p.Account.Id()
}
