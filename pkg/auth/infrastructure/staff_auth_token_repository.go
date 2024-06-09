package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/namhq1989/vocab-booster-server-admin/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-server-admin/core/error"
	"github.com/namhq1989/vocab-booster-server-admin/internal/database"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth/domain"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth/infrastructure/dbmodel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StaffAuthTokenRepository struct {
	db             *database.Database
	collectionName string
}

func NewStaffAuthTokenRepository(db *database.Database) StaffAuthTokenRepository {
	r := StaffAuthTokenRepository{
		db:             db,
		collectionName: database.Collections.StaffAuthToken,
	}
	r.ensureIndexes()
	return r
}

func (r StaffAuthTokenRepository) ensureIndexes() {
	var (
		ctx     = context.Background()
		opts    = options.CreateIndexes().SetMaxTime(time.Minute * 30)
		indexes = []mongo.IndexModel{
			{
				Keys: bson.D{{Key: "staffId", Value: 1}, {Key: "refreshToken", Value: 1}},
			},
		}
	)

	if _, err := r.collection().Indexes().CreateMany(ctx, indexes, opts); err != nil {
		fmt.Printf("index collection %s err: %v \n", r.collectionName, err)
	}
}

func (r StaffAuthTokenRepository) collection() *mongo.Collection {
	return r.db.GetCollection(r.collectionName)
}

func (r StaffAuthTokenRepository) CreateAuthToken(ctx *appcontext.AppContext, token domain.AuthToken) error {
	doc, err := dbmodel.AuthToken{}.FromDomain(token)
	if err != nil {
		return err
	}

	_, err = r.collection().InsertOne(ctx.Context(), &doc)
	return err
}

func (r StaffAuthTokenRepository) DeleteAuthToken(ctx *appcontext.AppContext, tokenID string) error {
	id, err := database.ObjectIDFromString(tokenID)
	if err != nil {
		return apperrors.Common.InvalidID
	}

	_, err = r.collection().DeleteOne(ctx.Context(), bson.M{
		"_id": id,
	})
	return err
}

func (r StaffAuthTokenRepository) FindAuthToken(ctx *appcontext.AppContext, refreshToken string) (*domain.AuthToken, error) {
	var doc dbmodel.AuthToken
	err := r.collection().FindOne(ctx.Context(), bson.M{
		"refreshToken": refreshToken,
	}).Decode(&doc)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	} else if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, apperrors.Auth.InvalidAuthToken
	}

	// respond
	result := doc.ToDomain()
	return &result, nil
}
