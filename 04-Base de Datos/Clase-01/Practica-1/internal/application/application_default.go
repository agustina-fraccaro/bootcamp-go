package application

import (
	"app/internal/handler"
	"app/internal/repository"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-sql-driver/mysql"
)

type ConfigDefault struct {
	// Database is the database configuration
	Database mysql.Config
	// Address is the address of the application
	Address string
}

// NewApplicationDefault creates a new default application.
func NewApplicationDefault(cfg *ConfigDefault) (a *ApplicationDefault) {
	// default
	cfgDefault := &ConfigDefault{
		Address: ":8080",
	}
	if cfg != nil {
		cfgDefault.Database = cfg.Database
		if cfg.Address != "" {
			cfgDefault.Address = cfg.Address
		}
	}

	return &ApplicationDefault{
		cfgDb: cfgDefault.Database,
		addr:  cfgDefault.Address,
	}
}

// ApplicationDefault is the default application.
type ApplicationDefault struct {
	// cfgDb is the database configuration
	cfgDb mysql.Config
	// addr is the address of the application
	addr string
}

// TearDown tears down the application.
func (a *ApplicationDefault) TearDown() (err error) {
	return
}

// SetUp sets up the application.
func (a *ApplicationDefault) SetUp() (err error) {
	// dependencies
	// - store
	//st := store.NewStoreProductJSON(a.filePathStore)
	// - repository
	//rp := repository.NewRepositoryProductStore(st)
	// - handler

	db, err := sql.Open("mysql", a.cfgDb.FormatDSN())
	if err != nil {
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return
	}

	// - repository: products
	rp := repository.NewRepositoryProductDB(db)
	rp_wh := repository.NewRepositoryWarehouseDB(db)
	hd := handler.NewHandlerProduct(rp)
	hd_wh := handler.NewHandlerWarehouse(rp_wh)
	rt := chi.NewRouter()
	// router
	// - middlewares
	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)
	// - endpoints
	rt.Route("/products", func(r chi.Router) {
		// GET /products/{id}
		r.Get("/{id}", hd.GetById())
		// POST /products
		r.Post("/", hd.Create())
		// PUT /products/{id}
		r.Put("/{id}", hd.UpdateOrCreate())
		// PATCH /products/{id}
		r.Patch("/{id}", hd.Update())
		// DELETE /products/{id}
		r.Delete("/{id}", hd.Delete())
		r.Get("/", hd.GetAll())
		r.Get("/warehouse/reportProducts", hd.GetByWarehouseId())
	})

	rt.Route("/warehouses", func(r chi.Router) {
		// GET /warehouses/{id}
		r.Get("/{id}", hd_wh.GetById())
		// POST /warehouses
		r.Post("/", hd_wh.Create())
		// GET /warehouses
		r.Get("/", hd_wh.GetAll())
	})

	err = http.ListenAndServe(a.addr, rt)
	if err != nil {
		return
	}

	return
}

// Run runs the application.
func (a *ApplicationDefault) Run() (err error) {
	//err = http.ListenAndServe(a.addr, rt)
	return
}
