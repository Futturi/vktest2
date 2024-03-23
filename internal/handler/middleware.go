package handler

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strings"
)

const (
	userHeader = "userId"
)

func (h *Handler) CheckIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": "your auth header is empty",
		})
		slog.Info("auth header problem")
		return
	}
	headerPart := strings.Split(header, " ")
	if len(headerPart) != 2 {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": "incorrect header",
		})
		slog.Info("incorrect header")
		return
	}
	id, err := h.service.SetHeader(headerPart[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "internal server error",
		})
		slog.Error("server error with middleware", slog.Any("error", err))
		return
	}
	c.Set(userHeader, id)
}
