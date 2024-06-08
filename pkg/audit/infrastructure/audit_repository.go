package infrastructure

import (
	"context"
	"fmt"
	"time"

	"github.com/namhq1989/vocab-booster-server-admin/core/appcontext"
	"github.com/namhq1989/vocab-booster-server-admin/internal/database"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/audit/domain"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/audit/infrastructure/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AuditRepository struct {
	db             *database.Database
	collectionName string
}

func NewAuditRepository(db *database.Database) AuditRepository {
	r := AuditRepository{
		db:             db,
		collectionName: database.Collections.Audit,
	}
	r.ensureIndexes()
	return r
}

func (r AuditRepository) ensureIndexes() {
	var (
		ctx     = context.Background()
		opts    = options.CreateIndexes().SetMaxTime(time.Minute * 30)
		indexes = []mongo.IndexModel{
			{
				Keys: bson.D{{Key: "entity._id", Value: 1}, {Key: "createdAt", Value: -1}},
			},
		}
	)

	if _, err := r.collection().Indexes().CreateMany(ctx, indexes, opts); err != nil {
		fmt.Printf("index collection %s err: %v \n", r.collectionName, err)
	}
}

func (r AuditRepository) collection() *mongo.Collection {
	return r.db.GetCollection(r.collectionName)
}

func (r AuditRepository) CreateAudit(ctx *appcontext.AppContext, audit domain.Audit) error {
	doc, err := model.Audit{}.FromDomain(audit)
	if err != nil {
		return err
	}

	_, err = r.collection().InsertOne(ctx.Context(), &doc)
	return err
}
