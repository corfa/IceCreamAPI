package handler

import (
	"github.com/gin-gonic/gin"
	"iceCreamApiWithDI/service"
)

type Handler struct {
	service service.IServices
}

func (h *Handler) InitHandler() *gin.Engine {

	router := gin.New()
	welcome := router.Group("/")
	{
		welcome.GET("/", h.Welcome)

	}
	auth := router.Group("/auth")
	{
		auth.POST("/sing-up", h.SingUp)
		auth.POST("/sing-in", h.SingIn)
	}

	return router
}

func NewHandler(service service.IServices) *Handler {
	return &Handler{service}
}
