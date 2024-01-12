package application

import (
	"app/internal/product"
	"app/internal/product/handlers"
	"app/internal/product/repository"
	"app/internal/product/service"
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
	rp := repository.NewProductMap(make(map[int]product.Product), 0)

	sv := service.NewProductDefault(rp)

	hd := handlers.NewDefaultProducts(sv)

	rt := chi.NewRouter()

	rt.Route("/products", func(rt chi.Router) {
		rt.Get("/{id}", hd.GetByID())
		rt.Get("/", hd.GetAll())
		rt.Post("/", hd.Create())
		rt.Put("/{id}", hd.Update())
		rt.Patch("/{id}", hd.UpdatePartial())
		rt.Delete("/{id}", hd.Delete())
	})

	err = http.ListenAndServe(h.addr, rt)
	return
}
