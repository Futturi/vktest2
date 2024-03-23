package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vktest2/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitHandlers() http.Handler {
	han := gin.Default()
	auth := han.Group("/auth")
	{
		auth.POST("/signup", h.SignUp)
		auth.POST("/signin", h.SignIn)
	}
	api := han.Group("/api")
	{
		api.POST("/", h.CheckIdentity, h.CreateAnn)
	}
	return han
}
