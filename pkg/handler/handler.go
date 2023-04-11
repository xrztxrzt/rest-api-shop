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

func (h *Handler) InitRoutes() *gin.Engine { //инициализация всех эндпоинтов
	router := gin.New()

	api := router.Group("/api")
	{
		product := api.Group("/product")
		{
			product.POST("/", h.createProduct)
			product.GET("/", h.getAllProduct)
			product.GET("/:id", h.getProductById)
			product.PUT("/:id", h.updateProduct)
			product.DELETE("/:id", h.deleteProduct)
		}
	}

	return router
}
