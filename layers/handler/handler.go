package handler

import (
	"github.com/gin-gonic/gin"
	"iceCreamApiWithDI/layers/service"
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
	iceCream:=router.Group("/ice-cream")
	{
		iceCream.POST("/add",h.addIceCream)
		iceCream.POST("/del",h.DelIceCream)

	}

	return router
}

func NewHandler(service service.IServices) *Handler {
	return &Handler{service}
}
