package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) Welcome(c *gin.Context) {
	c.JSON(200, gin.H{
		"welcome": "it is api fro ice cream store",
	})

}
