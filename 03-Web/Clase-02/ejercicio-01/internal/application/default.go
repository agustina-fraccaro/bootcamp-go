package application

import (
	"app/internal"
	"app/internal/handler"
	"app/internal/repository"
	"app/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewDefaultHTTP(addr string) *DefaultHTTP {

	return &DefaultHTTP{
		addr: addr,
	}
}

type DefaultHTTP struct {
	addr string
}

func (h *DefaultHTTP) Run() (err error) {
	// - repository
	rp := repository.NewProductMap(make(map[int]internal.Product), 0)
	// - service
	sv := service.NewProductDefault(rp)
	// - handler
	hd := handler.NewDefaultProducts(sv)
	// - router
	rt := chi.NewRouter()
	//   endpoints
	rt.Post("/products", hd.Create())

	// run http server
	err = http.ListenAndServe(h.addr, rt)
	return
}
