package entity

import "fmt"

type Entity interface {
	Id() string
	TypeName() string
}

func Key(e Entity) string {
	return fmt.Sprintf("%s:%s", e.TypeName(), e.Id())
}

func String(e Entity) string {
	return fmt.Sprintf("[%s(Id=%s)]", e.TypeName(), e.Id())
}
