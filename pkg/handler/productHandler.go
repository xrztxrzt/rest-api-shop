package handler

import (
	"net/http"
	"rest-api/pkg/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createProduct(c *gin.Context) {
	var input models.Product
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	product, err := h.services.ProductList.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})

}

func (h *Handler) getAllProduct(c *gin.Context) {

}

func (h *Handler) getSingleProduct(c *gin.Context) {

}

func (h *Handler) updateProduct(c *gin.Context) {

}

func (h *Handler) deleteProduct(c *gin.Context) {

}
