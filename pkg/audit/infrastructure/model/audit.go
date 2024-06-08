package model

import (
	"time"

	apperrors "github.com/namhq1989/vocab-booster-server-admin/core/error"
	"github.com/namhq1989/vocab-booster-server-admin/internal/database"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/audit/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Audit struct {
	ID        primitive.ObjectID `bson:"_id"`
	ActorID   primitive.ObjectID `bson:"actorId"`
	Action    string             `bson:"action"`
	Entity    Entity             `bson:"entity"`
	SourceIp  string             `bson:"sourceIp"`
	CreatedAt time.Time          `bson:"createdAt"`
}

type Entity struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
}

func (m Audit) ToDomain() domain.Audit {
	return domain.Audit{
		ID:      m.ID.Hex(),
		Action:  domain.ToAction(m.Action),
		ActorID: m.ActorID.Hex(),
		Entity: domain.Entity{
			ID:   m.Entity.ID,
			Name: m.Entity.Name,
		},
		SourceIp:  m.SourceIp,
		CreatedAt: m.CreatedAt,
	}
}

func (Audit) FromDomain(audit domain.Audit) (*Audit, error) {
	id, err := database.ObjectIDFromString(audit.ID)
	if err != nil {
		return nil, apperrors.Common.InvalidID
	}

	aid, err := database.ObjectIDFromString(audit.ActorID)
	if err != nil {
		return nil, apperrors.Common.InvalidID
	}

	return &Audit{
		ID:        id,
		Action:    audit.Action.String(),
		ActorID:   aid,
		Entity:    Entity{ID: audit.Entity.ID, Name: audit.Entity.Name},
		SourceIp:  audit.SourceIp,
		CreatedAt: audit.CreatedAt,
	}, nil
}
