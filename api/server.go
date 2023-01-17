package api

import (
	"fmt"
	db "github/promo-news-app/db/sqlc"
	"github/promo-news-app/token"
	"github/promo-news-app/util"

	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	config     util.Config
	store      db.DBStore
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config, store db.DBStore) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.JWTSecretKey)
	if err != nil {
		return nil, fmt.Errorf("Cannot Create Token Maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.routeConfig()
	return server, nil
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) routeConfig() {
	router := gin.Default()
	router.GET("/promo/:id", server.getPromo)
	router.GET("/promo", server.getListPromo)

	router.GET("/store/:id", server.getStore)
	router.GET("/store", server.getListStore)

	router.POST("/user", server.createUser)
	router.POST("/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/promo", server.createPromo)

	authRoutes.GET("/promo/toogle/:id", server.tooglePromo)
	authRoutes.PUT("/promo", server.updatePromo)

	authRoutes.POST("/store", server.createStore)

	authRoutes.GET("/store/toogle/:id", server.toogleStore)
	authRoutes.PUT("/store", server.updateStore)

	authRoutes.POST("/operational-time", server.createOperationalTime)
	authRoutes.GET("/operational-time/:id", server.getOperationalTime)
	authRoutes.GET("/operational-time", server.getListOperationalTime)
	authRoutes.GET("/operational-time/toogle/:id", server.toogleOperationalTime)
	authRoutes.PUT("/operational-time", server.updateOperationalTime)

	authRoutes.GET("/user/:id", server.getUser)
	authRoutes.GET("/user", server.getListUser)
	authRoutes.GET("/user/toogle/:id", server.toogleUser)
	authRoutes.PUT("/user", server.updateUser)

	server.router = router
}
