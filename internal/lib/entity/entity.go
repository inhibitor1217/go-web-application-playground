package entity

import "fmt"

type Entity interface {
	Identifier() string
	TypeName() string
}

func Key(e Entity) string {
	return fmt.Sprintf("%s:%s", e.TypeName(), e.Identifier())
}

func String(e Entity) string {
	return fmt.Sprintf("[%s(Id=%s)]", e.TypeName(), e.Identifier())
}
