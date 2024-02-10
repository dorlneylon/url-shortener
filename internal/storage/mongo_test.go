package storage

import (
	"testing"
	"time"
	"url-shortener/internal/config"
	"url-shortener/internal/models"
)

func TestMongo(t *testing.T) {
	cfg := config.Config{
		Mongo: config.Mongo{
			Database: "url-shortener",
			URI:      "mongodb://127.0.0.1:27017",
			Username: "admin",
			Password: "admin",
		},
		JwtSecret: "secret",
		Host:      "127.0.0.1",
		Port:      "3000",
	}
	storage := NewMongo(&cfg)

	_, err := storage.DeleteUser("test")

	if err != nil {
		t.Fatal(err)
	}

	_, err = storage.InsertUser(models.User{
		Name:     "test",
		Password: "test",
	})

	if err != nil {
		t.Fatal(err)
	}

	user, err := storage.GetUser("test")

	if err != nil {
		t.Fatal(err)
	}

	t.Log(user)
}

func TestMongoShortenings(t *testing.T) {
	cfg := config.Config{
		Mongo: config.Mongo{
			Database: "url-shortener",
			URI:      "mongodb://127.0.0.1:27017",
			Username: "admin",
			Password: "admin",
		},
		JwtSecret: "secret",
		Host:      "127.0.0.1",
		Port:      "3000",
	}

	link := models.Shortening{
		Alias:  "test",
		Author: "test",
		Url:    "test",
		Clicks: 0,
	}

	storage := NewMongo(&cfg)
	_, err := storage.InsertShortening(link)

	if err != nil {
		t.Fatal(err)
	}

	short, err := storage.GetByShortened(link.Alias)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(short)
}

func TestMongoUsers(t *testing.T) {
	cfg := config.Config{
		Mongo: config.Mongo{
			Database: "url-shortener",
			URI:      "mongodb://127.0.0.1:27017",
			Username: "admin",
			Password: "admin",
		},
		JwtSecret: "secret",
		Host:      "127.0.0.1",
		Port:      "3000",
	}

	storage := NewMongo(&cfg)

	_, _ = storage.DeleteUser("test")

	_, err := storage.InsertUser(models.User{
		Name:     "test",
		Password: "test",
	})

	if err != nil {
		t.Fatal(err)
	}

	user, err := storage.GetUser("test")

	if err != nil {
		t.Fatal(err)
	}

	t.Log(user)
}

func TestMongoIncrements(t *testing.T) {
	cfg := config.Config{
		Mongo: config.Mongo{
			Database: "url-shortener",
			URI:      "mongodb://127.0.0.1:27017",
			Username: "admin",
			Password: "admin",
		},
		JwtSecret: "secret",
		Host:      "127.0.0.1",
		Port:      "3000",
	}
	storage := NewMongo(&cfg)
	link := models.Shortening{
		Alias:     "test",
		Url:       "test",
		Author:    "test",
		Clicks:    0,
		Timestamp: time.Now().Unix(),
	}

	_, _ = storage.DeleteShortening(link.Alias)
	_, err := storage.InsertShortening(link)
	if err != nil {
		t.Fatal(err)
	}

	_, err = storage.IncrementClicks(link.Alias)
	if err != nil {
		t.Fatal(err)
	}

	ln, err := storage.GetByShortened(link.Alias)
	if ln.(models.Shortening).Clicks != 1 {
		t.Fatal("Increments work incorrectly")
	}
}
