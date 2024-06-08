package domain

import (
	"time"

	"github.com/namhq1989/vocab-booster-server-admin/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-server-admin/core/error"
	"github.com/namhq1989/vocab-booster-server-admin/internal/database"
)

type AuditRepository interface {
	CreateAudit(appCtx *appcontext.AppContext, audit Audit) error
}

type Audit struct {
	ID        string
	Action    Action
	ActorID   string
	Entity    Entity
	SourceIp  string
	CreatedAt time.Time
}

type Entity struct {
	ID   string
	Name string // table name
}

func NewAudit(action, actorID, entityID, entityName, sourceIp string) (*Audit, error) {
	if actorID == "" {
		return nil, apperrors.Audit.InvalidActor
	}

	if entityID == "" || entityName == "" {
		return nil, apperrors.Audit.InvalidEntity
	}

	dAction := ToAction(action)
	if !dAction.IsValid() {
		return nil, apperrors.Audit.InvalidAction
	}

	return &Audit{
		ID:      database.NewStringID(),
		Action:  dAction,
		ActorID: actorID,
		Entity: Entity{
			ID:   entityID,
			Name: entityName,
		},
		SourceIp:  sourceIp,
		CreatedAt: time.Now(),
	}, nil
}
