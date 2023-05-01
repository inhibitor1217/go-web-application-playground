package auth

type Principal interface {
	Type() string
	Id() string
}
