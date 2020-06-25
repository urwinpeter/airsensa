package requests

import (
	"log"
	"net/http"
)

type Handler struct {
	host, port string // TODO: make config file for these
}

func NewHandler(host, port string) *Handler {
	return &Handler{host, port}
}

func (h *Handler) LoadHandler(f func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc("/", f)
	err := http.ListenAndServe(h.host+":"+h.port, nil)
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}
