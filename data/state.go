package data

import (
	"database/sql"
)

type AppState struct {
	Db *sql.DB
}
