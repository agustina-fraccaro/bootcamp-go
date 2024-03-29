package application

import (
	"app/internal/auth"
	"app/internal/product/handlers"
	"app/internal/product/repository"
	"app/internal/product/service"
	"app/internal/product/storage"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewDefaultHTTP(addr string, token string, path string) *DefaultHTTP {

	defaultCfg := &DefaultHTTP{
		addr:  ":8080",
		token: "",
		path:  "",
	}

	if addr != "" {
		defaultCfg.addr = addr
	}
	if token != "" {
		defaultCfg.token = token
	}
	if path != "" {
		defaultCfg.path = path
	}

	return &DefaultHTTP{
		addr:  defaultCfg.addr,
		token: defaultCfg.token,
		path:  defaultCfg.path,
	}
}

type DefaultHTTP struct {
	addr  string
	token string
	path  string
}

func (h *DefaultHTTP) Run() (err error) {
	auth := auth.NewAuthTokenBasic(h.token)

	st := storage.NewStorageProductJSON(h.path)

	rp := repository.NewRepositoryProductStore(*st, 500)

	sv := service.NewProductDefault(rp)

	hd := handlers.NewDefaultProducts(sv, auth)

	rt := chi.NewRouter()

	rt.Route("/products", func(rt chi.Router) {
		rt.Get("/{id}", hd.GetByID())
		rt.Get("/", hd.GetAll())
		rt.Post("/", hd.Create())
		rt.Put("/{id}", hd.Update())
		rt.Patch("/{id}", hd.UpdatePartial())
		rt.Delete("/{id}", hd.Delete())
		rt.Get("/consumer-price", hd.GetConsumerPrice())
	})

	err = http.ListenAndServe(h.addr, rt)
	return
}
