package handler

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"vktest2/internal/models"
)

// @Summary SingUp
// @Tags auth
// @Description create account 4 user
// @ID create-account-user
// @Accept json
// @Produce json
// @Param input body models.User true "account info"
// @Success 200 {object} models.User
// @Failure default {string} error
// @Router /auth/signup [post]
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
		} else if err.Error() == "your username is incorrect, no need to put specific symphols" ||
			err.Error() == "your password mush have 1 upper symphol" || err.Error() == "your password mush have 1 of specific elements: @!$#%^*&-+=/;" {
			c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "problem with server",
		})
		slog.Error("error with server", slog.Any("error", err))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":       id,
		"username": user.Username,
		"password": user.Password,
	})
	slog.Info("create user", slog.Int("id", id))
}

// @Summary SignIn
// @Tags auth
// @Description login to account
// @ID login-account-user
// @Accept json
// @Produce json
// @Param input body models.User true "account info"
// @Success 200 {string} token
// @Failure default {string} error
// @Router /auth/signin [post]
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
