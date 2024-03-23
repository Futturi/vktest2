package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"vktest2/internal/models"
)

// @Summary CreateAnn
// @Security ApiKeyAuth
// @Tags announcements
// @Description insert announcement
// @ID insert-announcement
// @Accept json
// @Produce json
// @Success 200 {object} models.Annunc
// @Failure default {string} error
// @Router /api/ [post]
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
	if !exists || idUser == "false" {
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

// @Summary GetAnnouncements
// @Security ApiKeyAuth
// @Tags announcements
// @Description get announcement
// @ID get-announcement
// @Accept json
// @Produce json
// @Success 200 {object} []models.AnnuncARes
// @Failure default {string} error
// @Param   sort     query    string     false        "data or price, need to know what sort to use"
// @Param   sortTo     query    string     false        "up/down, need to know how to sort values"
// @Param   minprice     query    string     false        "value of min price"
// @Param   maxprice     query    string     false        "value of max price"
// @Router /api/{page} [get]
func (h *Handler) GetAnns(c *gin.Context) {
	page := c.Param("page")
	sor := c.Query("sort")
	sorTo := c.Query("sortto")
	minP := c.Query("minprice")
	maxP := c.Query("maxprice")
	idUser, exist := c.Get(userHeader)
	if !exist || idUser == "false" {
		ann, err := h.service.GetAnnsWithoutAuth(page, sor, sorTo, minP, maxP)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
				"error": "error with server",
			})
			slog.Error("error", slog.Any("error", err))
			return
		}
		c.JSON(http.StatusOK, ann)
	} else {
		ann, err := h.service.GetAnns(page, sor, sorTo, minP, maxP, fmt.Sprint(idUser))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
				"error": "error with server",
			})
			slog.Error("error", slog.Any("error", err))
			return
		}
		c.JSON(http.StatusOK, ann)
	}
}
