package requests

import (
	"fmt"
	"log"
	"net/http"

	"github.com/urwinpeter/airsensa/service"
)

type Handler struct {
	host, port  string // TODO: make config file for these
	dataservice *service.DataService
}

func NewHandler(host, port string, service *service.DataService) *Handler {
	return &Handler{host, port, service}
}

func (h *Handler) onRequest(w http.ResponseWriter, r *http.Request) {
	h.display("Pollution", w, r)
}

func (h *Handler) onShoesRequest(w http.ResponseWriter, r *http.Request) {
	h.display("Shoes", w, r)
}

func (h *Handler) display(key string, w http.ResponseWriter, r *http.Request) {
	data, found := h.dataservice.GetFromCache(key)
	if found {
		log.Print(
			"Key Found in Cache with value as :: ",
			data.(string),
		)
		fmt.Fprintf(w, data.(string))
	} else {
		log.Print("Key Not Found in Cache :: ", key)
		fmt.Fprintf(w, "Key Not Found in Cache")
	}
}

func (h *Handler) LoadHandler() {
	http.HandleFunc("/", h.onRequest)
	http.HandleFunc("/Shoes", h.onShoesRequest)
	err := http.ListenAndServe(h.host+":"+h.port, nil)
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}

}
