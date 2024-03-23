package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"vktest2/internal/models"
)

func (h *Handler) CreateAnn(c *gin.Context) {
	var ann models.Annunc
	if err := c.BindJSON(&ann); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": "incorrect request body",
		})
		slog.Info("user error with body", slog.Any("error", err))
		return
	}

	idUser, exists := c.Get(userHeader)
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"error": "unautorized",
		})
		slog.Info("user trying to createann without auth")
		return
	}
	id, err := h.service.CreateAnn(ann, fmt.Sprint(idUser))
	if err != nil {
		if err.Error() == "too big body" || err.Error() == "incorrect name" || err.Error() == "incorrect price" {
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
				"error": "incorrect request",
			})
			slog.Info("error with request", slog.Any("error", err))
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "server error",
		})
		slog.Error("error with server", slog.Any("error", err))
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id":    id,
		"name":  ann.Name,
		"body":  ann.Body,
		"image": ann.Image,
		"price": ann.Price,
	})
}
