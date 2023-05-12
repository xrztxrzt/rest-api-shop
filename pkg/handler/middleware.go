package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	roleCtx             = "userRole"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		newErrorResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	userId, roleId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
	c.Set(roleCtx, roleId)

	switch c.Request.Method {

	}
}

func (h *Handler) authorization(c *gin.Context) {
	userRoleStr, err := getUserRole(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userRole, err := strconv.Atoi(userRoleStr)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	switch c.Request.Method {
	case http.MethodGet:
		//Only users with role "admin(id=1)" or "user(id=2)" can perform a get request

		if userRole != 1 && userRole != 2 {
			newErrorResponse(c, http.StatusForbidden, "Access closed")
			return
		}
	case http.MethodPost:
		if userRole != 1 {
			newErrorResponse(c, http.StatusForbidden, "Access closed")
			return
		}

	case http.MethodPut:
		if userRole != 1 {
			newErrorResponse(c, http.StatusForbidden, "Access closed")
			return
		}

	case http.MethodDelete:
		if userRole != 1 {
			newErrorResponse(c, http.StatusForbidden, "Access closed")
			return
		}

	default:
		// If the HTTP method is not supported, return an error
		newErrorResponse(c, http.StatusBadRequest, "invalid HTTP method")
		return
	}

	c.Next()
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id not found")
	}
	return idInt, nil
}

func getUserRole(c *gin.Context) (string, error) {
	role, ok := c.Get(roleCtx)
	if !ok {
		return "", errors.New("user role not found")
	}
	roleStr, ok := role.(string)
	if !ok {
		return "", errors.New("user role not found")
	}
	return roleStr, nil
}
