package lifecycle

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/urwinpeter/airsensa/requests"
	"github.com/urwinpeter/airsensa/service"
	"github.com/urwinpeter/airsensa/storage"
)

type lifecycle struct {
	dataservice *service.DataService
	handler     *requests.Handler
}

func NewLifecycle(dbconn *sql.DB) *lifecycle {
	service := service.NewDataService(dbconn)
	handler := requests.NewHandler("localhost", "8080")
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

func (lc *lifecycle) loadCache(data []storage.Datum) {
	lc.dataservice.LoadCache(data)
}

func (lc *lifecycle) loadRequestHandler() {
	lc.handler.LoadHandler(lc.onRequest)
}

func (lc *lifecycle) onRequest(w http.ResponseWriter, r *http.Request) {
	lc.dataservice.GetFromCache(w, r)
}
