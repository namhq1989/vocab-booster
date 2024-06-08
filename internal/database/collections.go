package database

import "go.mongodb.org/mongo-driver/mongo"

var Collections = struct {
	Audit string

	User string

	Staff          string
	StaffAuthToken string
}{
	Audit: "audit.audits",

	User: "user.users",

	Staff:          "staff.staffs",
	StaffAuthToken: "staff.authTokens",
}

func (db Database) GetCollection(table string) *mongo.Collection {
	return db.mongo.Collection(table)
}
