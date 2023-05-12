package handler

import (
	"rest-api/pkg/service"

	"github.com/gin-gonic/gin"
)

// внедрение зависимостей
type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/create-role", h.createRole)
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)

	}

	product := router.Group("/product")
	{
		product.GET("/", h.userIdentity, h.authorization, h.getAllProduct)
		product.GET("/:id", h.userIdentity, h.authorization, h.getProductById)
		product.POST("/", h.userIdentity, h.authorization, h.createProduct)
		product.PUT("/:id", h.userIdentity, h.authorization, h.updateProduct)
		product.DELETE("/:id", h.userIdentity, h.authorization, h.deleteProduct)
	}

	return router
}
