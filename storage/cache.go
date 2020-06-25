package storage

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/patrickmn/go-cache"
)

type Cache struct {
	cache *cache.Cache
}

func NewCache() *Cache {
	return &Cache{cache.New(0, 0)}
}

func (c *Cache) LoadData(items []Datum) {
	for _, item := range items {
		c.cache.Set(item.Category, item.Name, cache.DefaultExpiration)
	}
}

func (c *Cache) GetData(w http.ResponseWriter, r *http.Request) {
	urlPathElements := strings.Split(r.URL.Path, "/")
	foo, found := c.cache.Get(urlPathElements[1])
	if found {
		log.Print(
			"Key Found in Cache with value as :: ",
			foo.(string),
		)
		fmt.Fprintf(w, "Hello "+foo.(string))
	} else {
		log.Print("Key Not Found in Cache :: ", "foo")
		fmt.Fprintf(w, "Key Not Found in Cache")
	}
}
