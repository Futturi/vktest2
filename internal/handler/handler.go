package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	_ "vktest2/docs"
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
	han.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth := han.Group("/auth")
	{
		auth.POST("/signup", h.SignUp)
		auth.POST("/signin", h.SignIn)
	}
	api := han.Group("/api", h.CheckIdentity)
	{
		api.POST("/", h.CreateAnn)
		api.GET("/:page", h.GetAnns)
	}
	return han
}
