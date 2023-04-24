CREATE TABLE accounts (
  id TEXT PRIMARY KEY,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
  email TEXT NOT NULL,
  password_hash TEXT NOT NULL,
  display_name TEXT,
  touched_at TIMESTAMP WITHOUT TIME ZONE
);

CREATE UNIQUE INDEX index_accounts_on_email ON accounts (email);
