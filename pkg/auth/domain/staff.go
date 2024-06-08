package domain

import (
	"github.com/namhq1989/vocab-booster-server-admin/core/appcontext"
)

type StaffHub interface {
	FindOneByID(ctx *appcontext.AppContext, id string) (*Staff, error)
	FindOneByEmail(ctx *appcontext.AppContext, email string) (*Staff, error)
}

type Staff struct {
	ID      string
	Name    string
	Email   string
	IsAdmin bool
}
