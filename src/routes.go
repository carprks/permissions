package src

import (
	"fmt"
	"github.com/carprks/permissions/src/permissions"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/keloran/go-healthcheck"
	"github.com/keloran/go-probe"
	"os"
	"time"
)

func Routes() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	// Health check
	router.Get(fmt.Sprintf("%s/healthcheck", os.Getenv("SITE_PREFIX")), healthcheck.HTTP)

	// Probe
	router.Get(fmt.Sprintf("%s/probe", os.Getenv("SITE_PREFIX")), probe.HTTP)
	router.Get("/probe", probe.HTTP)

	// Create
	router.Post(fmt.Sprintf("%s/", os.Getenv("SITE_PREFIX")), permissions.Create)

	// Retrieve
	// router.Get(fmt.Sprintf("%s/", os.Getenv("SITE_PREFIX")), permissions.RetrieveAll)

	// System
	// router.Route(fmt.Sprintf("%s/system/{permissionID}", os.Getenv("SITE_PREFIX")), func(r chi.Router) {
	// 	r.Get("/", permissions.RetrieveSystem)
	// 	r.Put("/", permissions.UpdateSystem)
	// 	r.Delete("/", permissions.DeleteSystem)
	// })

	// User
	// router.Route(fmt.Sprintf("%s/user/{permissionID}", os.Getenv("SITE_PREFIX")), func(r chi.Router) {
	// 	r.Get("/", permissions.RetrieveUser)
	// 	r.Put("/", permissions.UpdateUser)
	// 	r.Delete("/", permissions.DeleteUser)
	// })

	return router
}