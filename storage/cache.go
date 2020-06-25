package storage

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type Cache struct {
	cache *cache.Cache
}

func NewCache(expire, clean time.Duration) *Cache {
	return &Cache{cache.New(expire, clean)}
}

func (c *Cache) LoadData(items []byte) {
	c.cache.Set("Pollution", string(items), cache.DefaultExpiration)
	c.cache.Set("Shoes", "Trainers", cache.DefaultExpiration)
}

func (c *Cache) GetData(key string) (interface{}, bool) {
	return c.cache.Get(key)
}
