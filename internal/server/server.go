package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/regalangcom/go-shop-api/internal/config"
	"github.com/regalangcom/go-shop-api/internal/services"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Server struct {
	config *config.Config
	db     *gorm.DB
	logger *zerolog.Logger
	/*
		solution for the coupling issue is to inject the services into the server struct, so that the server can use the services without having to create them each time a request is made. This way, the server can be more easily tested and maintained, as the services can be mocked or replaced with different implementations if needed.
	*/
	authService    *services.AuthService
	productService *services.ProductService
	userService    *services.UserService
}

func New(cfg *config.Config,
	db *gorm.DB,
	logger *zerolog.Logger,
	authService *services.AuthService,
	productService *services.ProductService,
	userService *services.UserService,
) *Server {
	return &Server{
		config:         cfg,
		db:             db,
		logger:         logger,
		authService:    authService,
		productService: productService,
		userService:    userService,
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

	// auth route
	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", s.Register)
			auth.POST("/login", s.Login)
			auth.POST("/refresh", s.RefreshToken)
			auth.POST("/logout", s.Logout)
		}

		protected := api.Group("/")
		protected.Use(s.AuthMiddleware())
		{
			// protected routes go here
			users := protected.Group("/users")
			{
				userRoutes := users
				userRoutes.GET("/profile", s.GetProfile)
				userRoutes.PUT("/profile", s.UpdateProfile)
			}

			categories := protected.Group("/categories")
			{
				categoriesRoute := categories
				categoriesRoute.POST("/", s.AdminMiddleware(), s.createCategory)
				categoriesRoute.PUT("/:id", s.AdminMiddleware(), s.updateCategory)
				categoriesRoute.DELETE("/:id", s.AdminMiddleware(), s.deleteCategory)
			}

			products := protected.Group("/products")
			{
				productsRoute := products
				productsRoute.POST("/", s.AdminMiddleware(), s.createProduct)
				productsRoute.PUT("/:id", s.AdminMiddleware(), s.updateProduct)
				productsRoute.DELETE("/:id", s.AdminMiddleware(), s.deleteProduct)
			}

		}

		// public routes
		api.GET("/categories", s.getCategories)
		api.GET("/products", s.getProducts)
		api.GET("/products/:id", s.getProduct)

	}

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
