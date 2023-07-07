package apigomart

import (
	"fmt"
	"net/http"

	"github.com/DeneesK/go-mart/internal/api-go-mart/config"
	"github.com/DeneesK/go-mart/internal/api-go-mart/repository/pgrepo"
	"github.com/DeneesK/go-mart/internal/storage-go-mart"

	"github.com/DeneesK/go-mart/internal/api-go-mart/rest/routers"
	"github.com/DeneesK/go-mart/internal/logger"
	"github.com/go-chi/chi/v5"
)

func Run() {
	cfg := config.NewConfig()

	log, err := logger.LoggerInitializer(cfg.LogLevel)
	if err != nil {
		return
	}

	var (
		r   = chi.NewRouter()
		srv = &http.Server{
			Addr:    cfg.ServerAddress,
			Handler: r,
		}
	)

	pg, err := storagegomart.NewPostgresStorage(cfg.DSN)
	if err != nil {
		fmt.Println(err.Error())
	}

	repo := pgrepo.NewRepository(pg)

	routers.GetRoutes(r, repo, log)

	log.Infof("Server is running")
	if err = srv.ListenAndServe(); err != nil {
		//logging
		log.Fatal(err)
	}
}
