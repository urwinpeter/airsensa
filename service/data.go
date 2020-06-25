package service

import (
	"database/sql"
	"time"

	"github.com/urwinpeter/airsensa/storage"
)

type DataService struct {
	db    *storage.ItemsDB
	cache *storage.Cache
}

func NewDataService(dbconn *sql.DB) *DataService {
	return &DataService{
		storage.NewItemsDB(dbconn),
		storage.NewCache(0, 0),
	}
}

func (service *DataService) LoadCache(data []byte) {
	service.cache.LoadData(data)
}

func (service *DataService) GetFromCache(key string) (interface{}, bool) {
	return service.cache.GetData(key)
}

func (service *DataService) GetFromDB(now, past time.Time) []byte {
	return service.db.GetData(now, past)
}
