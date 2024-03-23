package handler

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"vktest2/internal/models"
)

func (h *Handler) SignUp(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid request",
		})
		slog.Info("error with marshalling body", slog.Any("error", err))
		return
	}
	id, err := h.service.SignUp(user)
	if err != nil {
		if err.Error() == "incorrect password" || err.Error() == "incorrect username" {
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
				"error": "incorrect request body",
			})
			slog.Info("error with request", slog.Any("error", err))
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "problem with server",
		})
		slog.Error("error with server", slog.Any("error", err))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
	slog.Info("create user", slog.Int("id", id))
}

func (h *Handler) SignIn(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid request",
		})
		slog.Info("error with marshalling body", slog.Any("error", err))
		return
	}
	token, err := h.service.SignIn(user)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
				"error": "account not finded in db",
			})
			slog.Info("error with request", slog.Any("error", err))
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "problem with server",
		})
		slog.Error("error with server", slog.Any("error", err))
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
	slog.Info("user logged", slog.String("token", token))
}
