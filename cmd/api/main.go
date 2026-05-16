package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/regalangcom/go-shop-api/internal/config"
	"github.com/regalangcom/go-shop-api/internal/database"
	"github.com/regalangcom/go-shop-api/internal/logger"
	"github.com/regalangcom/go-shop-api/internal/server"
)

func main() {

	log := logger.New()
	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load configuration")
	}

	db, err := database.New(&cfg.Database)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to database")
	}

	mainDB, err := db.DB()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get database instance")
	}

	// defer mainDB.Close()
	defer func() { _ = mainDB.Close() }()
	gin.SetMode(cfg.Server.GinDebug)

	srv := server.New(cfg, db, &log)
	router := srv.SetupRoute()

	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Server.Port),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Info().Msg(fmt.Sprintf("starting server on port %s", cfg.Server.Port))
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("failed to start server")
		}
	}()

	// create channel
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Info().Msg("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("failed to shutdown server")
	}

	log.Info().Msg(fmt.Sprintf("starting server on port %s", cfg.Server.Port))

}
