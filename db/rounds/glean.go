package rounds

import "database/sql"

//go:generate glean -table=rounds

type glean struct {
	Round
	Error sql.NullString
}
