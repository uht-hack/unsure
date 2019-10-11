package server

import (
	"github.com/uht-hack/unsure/db"
)

//go:generate genbackendsimpl
type Backends interface {
	UhtDB() *db.UhtDB
}
