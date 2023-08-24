package app

import (
	"avito_test/api"
	"avito_test/api/middleware"
	"avito_test/internals/cfg"
	"avito_test/internals/db"
	"avito_test/internals/handlers"
	"avito_test/internals/services"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog"
	// "github.com/sirupsen/logrus"
)

type AppServer struct {
	config cfg.Cfg
	ctx    context.Context
	srv    *http.Server
	db     *pgxpool.Pool
	logger *zerolog.Logger
}

func NewServer(config cfg.Cfg, cntx context.Context,logg *zerolog.Logger) *AppServer { //задаем поля нашего сервера, для его старта нам нужен контекст и конфигурация
	server := new(AppServer)
	server.ctx = cntx
	server.config = config
	server.logger = logg
	return server
}

func (server *AppServer) Serve() error {
	log.Println("Starting server")
	log.Println(server.config.GetDBString())
 
	var err error
	server.db, err = pgxpool.Connect(server.ctx, server.config.GetDBString())

	if err != nil {
		log.Fatalln(err)
	}

	storagePool := db.NewStorage(server.db)
	service := services.NewService(storagePool)
	orderHandler := handlers.NewHandler(service)

	routes := api.CreateRoutes(orderHandler) 
	routes.Use(middleware.RequestLog)                                                       

	server.srv = &http.Server{ 
		Addr:    "0.0.0.0:" + server.config.Port,
		Handler: routes,
	}

 

	server.logger.Info().Msg("Server started.")

	err = server.srv.ListenAndServe() 

	if err != nil && err != http.ErrServerClosed {
		server.logger.Err(err).Msg("Failure while serving")
		return err
	}

	return nil
}

func (server *AppServer) Shutdown() error {
	server.logger.Info().Msg("server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	server.db.Close() 

	defer func() {
		cancel()
	}()

	var err error

	if err = server.srv.Shutdown(ctxShutDown); err != nil { 
		server.logger.Err(err)

		err = fmt.Errorf("server shutdown failed %w. ", err)

		return err
	}
	server.logger.Info().Msg("Shutdown!")

	return nil
}