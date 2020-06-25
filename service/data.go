package service

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/urwinpeter/airsensa/storage"
)

type DataService struct {
	db    *storage.ItemsDB
	cache *storage.Cache
}

func NewDataService(dbconn *sql.DB) *DataService {
	itemsdb := storage.NewItemsDB(dbconn)
	cache := storage.NewCache()
	return &DataService{itemsdb, cache}
}

// Load 10 days of data into cache
func (service *DataService) LoadCache(data []storage.Datum) {
	service.cache.LoadData(data)
}

func (service *DataService) GetFromCache(w http.ResponseWriter, r *http.Request) {
	service.cache.GetData(w, r)
}

func (service *DataService) GetFromDB(now, past time.Time) []storage.Datum {
	items := service.db.GetData(now, past)
	return items
}
