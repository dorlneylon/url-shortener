package storage

import (
	"testing"
	"url-shortener/internal/config"
)

func TestNewMemcached(t *testing.T) {
	cfg := config.Config{
		Memcached: config.Memcached{
			URI: "127.0.0.1:11211",
		},
	}
	storage := NewMemcached(&cfg)
	defer storage.Disconnect()
	t.Log(storage)

	if err := storage.Set("test", "test"); err != nil {
		t.Fatal(err)
	}

	value, err := storage.Get("test")

	if err != nil {
		t.Fatal(err)
	}

	t.Log(value)

	if err := storage.Delete("test"); err != nil {
		t.Fatal(err)
	}
}
