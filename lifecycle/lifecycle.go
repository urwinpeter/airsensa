package lifecycle

import (
	"time"

	"github.com/urwinpeter/airsensa/service"
)

func Start() {
	service.Load()
}

func refresh() {
	service.Refresh()
	time.Sleep(10 * time.Second)
	refresh()
}
