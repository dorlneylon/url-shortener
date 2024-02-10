package storage

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"url-shortener/internal/config"
	"url-shortener/internal/models"
)

type Mongo struct {
	db *mongo.Database
}

func NewMongo(cfg *config.Config) *Mongo {
	addr := fmt.Sprintf(cfg.Mongo.URI)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(addr))
	if err != nil {
		panic(err)
	}
	return &Mongo{
		db: client.Database(cfg.Mongo.Database),
	}
}

func (m *Mongo) shortenings() *mongo.Collection {
	return m.db.Collection("shortenings")
}

func (m *Mongo) users() *mongo.Collection {
	return m.db.Collection("users")
}

func withTimeout(fn func(context.Context) (interface{}, error)) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return fn(ctx)
}

func (m *Mongo) Disconnect() (interface{}, error) {
	return withTimeout(func(ctx context.Context) (interface{}, error) {
		return true, m.db.Client().Disconnect(ctx)
	})
}

func (m *Mongo) InsertShortening(shortening models.Shortening) (interface{}, error) {
	return withTimeout(func(ctx context.Context) (interface{}, error) {
		_, err := m.shortenings().InsertOne(ctx, shortening)
		return true, err
	})
}

func (m *Mongo) GetByShortened(shortened string) (interface{}, error) {
	var short models.Shortening
	_, err := withTimeout(func(ctx context.Context) (interface{}, error) {
		return nil, m.shortenings().FindOne(ctx, bson.M{"alias": shortened}).Decode(&short)
	})
	return short, err
}

func (m *Mongo) DeleteShortening(shortened string) (interface{}, error) {
	return withTimeout(func(ctx context.Context) (interface{}, error) {
		_, err := m.shortenings().DeleteOne(ctx, bson.M{"alias": shortened})
		return true, err
	})
}

func (m *Mongo) IncrementClicks(shortened string) (interface{}, error) {
	return withTimeout(func(ctx context.Context) (interface{}, error) {
		_, err := m.shortenings().UpdateOne(
			ctx,
			bson.M{
				"alias": shortened,
			},
			bson.D{
				{
					"$inc",
					bson.D{
						{"clicks", 1},
					},
				},
			},
			options.Update().SetUpsert(true),
		)
		return true, err
	})
}

func (m *Mongo) InsertUser(user models.User) (interface{}, error) {
	return withTimeout(func(ctx context.Context) (interface{}, error) {
		_, err := m.users().InsertOne(ctx, user)
		return true, err
	})
}

func (m *Mongo) GetUser(name string) (interface{}, error) {
	return withTimeout(func(ctx context.Context) (interface{}, error) {
		var user models.User
		err := m.users().FindOne(ctx, bson.M{"name": name}).Decode(&user)
		return user, err
	})
}

func (m *Mongo) DeleteUser(name string) (interface{}, error) {
	return withTimeout(func(ctx context.Context) (interface{}, error) {
		del, err := m.users().DeleteOne(ctx, bson.M{"name": name})
		return del.DeletedCount, err
	})
}
