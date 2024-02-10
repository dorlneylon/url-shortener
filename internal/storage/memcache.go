package storage

import (
	"github.com/bradfitz/gomemcache/memcache"
	"url-shortener/internal/config"
)

type Memcached struct {
	client *memcache.Client
}

func NewMemcached(cfg *config.Config) *Memcached {
	return &Memcached{
		client: memcache.New(cfg.Memcached.URI),
	}
}

func (m *Memcached) Get(key string) (string, error) {
	item, err := m.client.Get(key)
	if err != nil {
		return "", err
	}
	return string(item.Value), nil
}

func (m *Memcached) Set(key, value string) error {
	return m.client.Set(&memcache.Item{Key: key, Value: []byte(value)})
}

func (m *Memcached) Delete(key string) error {
	return m.client.Delete(key)
}

func (m *Memcached) Ping() error {
	return m.client.Ping()
}

func (m *Memcached) Disconnect() error {
	return m.client.Close()
}
