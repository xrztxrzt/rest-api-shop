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
			cart.POST("/", h.createCart)
			cart.GET("/", h.getAllCart)
			cart.GET("/:id", h.getSingleCart)
			cart.PUT("/:id", h.updateCart)
			cart.DELETE("/id", h.deleteCart)
		}
		product := api.Group("/product")
		{
			product.POST("/", h.createProduct)
			product.GET("/", h.getAllProduct)
			product.GET("/:id", h.getSingleProduct)
			product.PUT("/:id", h.updateProduct)
			product.DELETE("/:id", h.deleteProduct)
		}
	}

	return router
}
