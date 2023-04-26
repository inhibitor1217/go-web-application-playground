package sqlschema

type Account struct {
	Id           string  `db:"id"`
	CreatedAt    string  `db:"created_at"`
	UpdatedAt    string  `db:"updated_at"`
	Email        string  `db:"email"`
	PasswordHash string  `db:"password_hash"`
	DisplayName  *string `db:"display_name"`
	TouchedAt    *string `db:"touched_at"`
}
