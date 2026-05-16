package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/regalangcom/go-shop-api/internal/config"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Server struct {
	config *config.Config
	db     *gorm.DB
	logger *zerolog.Logger
}

func New(cfg *config.Config, db *gorm.DB, logger *zerolog.Logger) *Server {
	return &Server{
		config: cfg,
		db:     db,
		logger: logger,
	}
}

func (s *Server) SetupRoute() *gin.Engine {
	r := gin.New()

	// add middleware, e.g. logging, recovery, CORS, etc.
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(s.corsMiddleware())

	// define routes
	r.GET("/health", s.healthCheck)

	return r
}

func (s *Server) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}

func (s *Server) corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
