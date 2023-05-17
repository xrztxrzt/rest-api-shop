package handler

import (
	"net/http"
	"rest-api/pkg/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createProduct(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input models.Product

	if err := c.BindJSON(&input); err != nil {
		h.logger.Errorf("Failed to bind JSON: %v", err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.ProductList.Create(userId, input)
	if err != nil {
		h.logger.Errorf("Failed to create product: %v", err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.logger.Infof("Product created with ID: %d", id)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllProductResponce struct {
	Data []models.Product `json:"data"`
}

func (h *Handler) getAllProduct(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	products, err := h.services.ProductList.GetAll(userId, models.Product{})
	if err != nil {
		h.logger.Errorf("Failed to get all products: %v", err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	h.logger.Infof("Retrieved %d products", len(products))

	c.JSON(http.StatusOK, getAllProductResponce{
		Data: products,
	})
}

func (h *Handler) getProductById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Errorf("Invalid ID parameter: %v", err)
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	product, err := h.services.GetById(userId, productId)
	if err != nil {
		h.logger.Errorf("Failed to get product by ID: %v", err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.logger.Infof("Retrieved product with ID: %d", productId)

	c.JSON(http.StatusOK, product)

}

func (h *Handler) updateProduct(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Errorf("Invalid ID parameter: %v", err)
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input models.UpdateProductInput
	if err := c.BindJSON(&input); err != nil {
		h.logger.Errorf("Failed to bind JSON: %v", err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Update(userId, productId, input); err != nil {
		h.logger.Errorf("Failed to update product: %v", err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.logger.Infof("Product updated with ID: %d", productId)
	c.JSON(http.StatusOK, statusResponce{"ok"})
}

func (h *Handler) deleteProduct(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		h.logger.Errorf("Failed to get user ID: %v", err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Errorf("Invalid ID parameter: %v", err)
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Delete(userId, productId)
	if err != nil {
		h.logger.Errorf("Failed to delete product:%v", err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.logger.Infof("Product deleted with ID: %d", productId)

	c.JSON(http.StatusOK, statusResponce{
		Status: "ok",
	})

}
