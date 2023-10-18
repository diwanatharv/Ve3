package Mongo

import (
	"awesomeProject/pkg/config"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongodbs struct {
	db *mongo.Collection
}

func MongoManager() *mongodbs {
	return &mongodbs{
		db: config.Makemongoserver().Collection("users"),
	}
}

type mongomethodsinterface interface {
	Insert(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	Deletekaro(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	Totalcount(ctx context.Context, opts ...*options.EstimatedDocumentCountOptions) (int64, error)
	Findusers(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (cur *mongo.Cursor, err error)
	Updateone(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	Findone(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
	Findanddelete(ctx context.Context, filter interface{}, opts ...*options.FindOneAndDeleteOptions) *mongo.SingleResult
}

func (m *mongodbs) Insert(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return m.db.InsertOne(ctx, document, opts...)
}
func (m *mongodbs) Deletekaro(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return m.db.DeleteOne(ctx, filter, opts...)
}
func (m *mongodbs) Totalcount(ctx context.Context, opts ...*options.EstimatedDocumentCountOptions) (int64, error) {
	return m.db.EstimatedDocumentCount(ctx, opts...)
}
func (m *mongodbs) Findusers(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (cur *mongo.Cursor, err error) {
	return m.db.Find(ctx, filter, opts...)
}
func (m *mongodbs) Updateone(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return m.db.UpdateOne(ctx, filter, update, opts...)
}
func (m *mongodbs) Findone(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	return m.db.FindOne(ctx, filter, opts...)
}
func (m *mongodbs) Findanddelete(ctx context.Context, filter interface{}, opts ...*options.FindOneAndDeleteOptions) *mongo.SingleResult {
	return m.db.FindOneAndDelete(ctx, filter, opts...)
}
