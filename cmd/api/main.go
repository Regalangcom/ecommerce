package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/regalangcom/go-shop-api/internal/config"
	"github.com/regalangcom/go-shop-api/internal/database"
	"github.com/regalangcom/go-shop-api/internal/logger"
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

	log.Info().Msg(fmt.Sprintf("starting server on port %s", cfg.Server.Port))

}
