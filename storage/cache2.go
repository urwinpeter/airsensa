package storage

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/patrickmn/go-cache"
)

// Move to config file
const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

var airCache *cache.Cache

func init() {
	airCache = cache.New(
		0, // Expiration Time
		0, // Clean Up
	)
	airCache.Set("meg", "lomax", cache.DefaultExpiration)
	airCache.Set("pete", "urwin", cache.DefaultExpiration)
}

func Load(category string, name string, price float32, datetime string) {
	airCache.Set(category, name, cache.DefaultExpiration)
}

func getFromCache(w http.ResponseWriter, r *http.Request) {
	urlPathElements := strings.Split(r.URL.Path, "/")
	foo, found := airCache.Get(urlPathElements[1])
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

func main() {
	http.HandleFunc("/", getFromCache)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}
