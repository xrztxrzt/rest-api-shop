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

	// userAuth := router.Group("/auth")
	// {
	// 	userAuth.POST("/sign-up") //регистрация
	// 	userAuth.POST("/sign-in") //авторизация
	// }
	api := router.Group("/api")
	{
		cart := api.Group("/cart")
		{
			cart.POST("/", h.createProduct)
			cart.GET("/", h.getAllProduct)
			cart.GET("/:id", h.getSingleProduct)
			cart.PUT("/:id", h.updateProduct)
			cart.DELETE("/id", h.deleteProduct)
		}
		product := api.Group("/product")
		{
			product.POST("/", h.createCart)
			product.GET("/", h.getAllCart)
			product.GET("/:id", h.getSingleCart)
			product.PUT("/:id", h.updateCart)
			product.DELETE("/:id", h.deleteCart)
		}
	}

	return router
}
