package lifecycle

import (
	"database/sql"
	"time"

	"github.com/urwinpeter/airsensa/requests"
	"github.com/urwinpeter/airsensa/service"
)

type lifecycle struct {
	dataservice *service.DataService
	handler     *requests.Handler
}

func NewLifecycle(dbconn *sql.DB) *lifecycle {
	service := service.NewDataService(dbconn)
	handler := requests.NewHandler("localhost", "8080", service)
	return &lifecycle{service, handler}
}

func (lc *lifecycle) Start() {
	now := time.Now()
	data := lc.dataservice.GetFromDB(
		now,
		now.Add(time.Hour*24*10*-1),
	)
	lc.loadCache(data)
	lc.loadRequestHandler()
}

func (lc *lifecycle) loadCache(data []byte) {
	lc.dataservice.LoadCache(data)
}

func (lc *lifecycle) loadRequestHandler() {
	lc.handler.LoadHandler()
}
