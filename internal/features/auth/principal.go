package auth

type Principal interface {
	Subject() string
}
