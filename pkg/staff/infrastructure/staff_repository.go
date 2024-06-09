package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/namhq1989/vocab-booster-server-admin/core/appcontext"
	"github.com/namhq1989/vocab-booster-server-admin/internal/database"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/staff/domain"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/staff/infrastructure/dbmodel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StaffRepository struct {
	db             *database.Database
	collectionName string
}

func NewStaffRepository(db *database.Database) StaffRepository {
	r := StaffRepository{
		db:             db,
		collectionName: database.Collections.Staff,
	}
	r.ensureIndexes()
	return r
}

func (r StaffRepository) ensureIndexes() {
	var (
		ctx     = context.Background()
		opts    = options.CreateIndexes().SetMaxTime(time.Minute * 30)
		indexes = []mongo.IndexModel{
			{
				Keys: bson.D{{Key: "email", Value: 1}, {Key: "status", Value: 1}, {Key: "role", Value: 1}, {Key: "createdAt", Value: -1}},
			},
		}
	)

	if _, err := r.collection().Indexes().CreateMany(ctx, indexes, opts); err != nil {
		fmt.Printf("index collection %s err: %v \n", r.collectionName, err)
	}
}

func (r StaffRepository) collection() *mongo.Collection {
	return r.db.GetCollection(r.collectionName)
}

func (r StaffRepository) CreateStaff(ctx *appcontext.AppContext, staff domain.Staff) error {
	doc, err := dbmodel.Staff{}.FromDomain(staff)
	if err != nil {
		return err
	}

	_, err = r.collection().InsertOne(ctx.Context(), &doc)
	return err
}

func (r StaffRepository) FindStaffByEmail(ctx *appcontext.AppContext, email string) (*domain.Staff, error) {
	// find
	var doc dbmodel.Staff
	if err := r.collection().FindOne(ctx.Context(), bson.M{
		"email": email,
	}).Decode(&doc); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	} else if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}

	// respond
	result := doc.ToDomain()
	return &result, nil
}

func (r StaffRepository) CountByEmail(ctx *appcontext.AppContext, email string) (int64, error) {
	return r.collection().CountDocuments(ctx.Context(), bson.M{
		"email": email,
	})
}
