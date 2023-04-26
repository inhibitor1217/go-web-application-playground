package sqlschema

import "time"

type Account struct {
	id           string  `db:"id"`
	createdAt    string  `db:"created_at"`
	updatedAt    string  `db:"updated_at"`
	email        string  `db:"email"`
	passwordHash string  `db:"password_hash"`
	displayName  *string `db:"display_name"`
	touchedAt    *string `db:"touched_at"`
}

func (a *Account) Id() string {
	return a.id
}

func (a *Account) CreatedAt() time.Time {
	createdAt, err := time.Parse(time.RFC3339, a.createdAt)
	if err != nil {
		// TODO better handling of this somehow?
		panic(err)
	}
	return createdAt
}

func (a *Account) UpdatedAt() time.Time {
	updatedAt, err := time.Parse(time.RFC3339, a.updatedAt)
	if err != nil {
		// TODO better handling of this somehow?
		panic(err)
	}
	return updatedAt
}

func (a *Account) Email() string {
	return a.email
}

func (a *Account) PasswordHash() string {
	return a.passwordHash
}

func (a *Account) DisplayName() *string {
	return a.displayName
}

func (a *Account) TouchedAt() *time.Time {
	if a.touchedAt == nil {
		return nil
	}
	touchedAt, err := time.Parse(time.RFC3339, *a.touchedAt)
	if err != nil {
		// TODO better handling of this somehow?
		panic(err)
	}
	return &touchedAt
}
