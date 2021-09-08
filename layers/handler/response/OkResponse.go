package response

import "github.com/gin-gonic/gin"

func OkResponse(c *gin.Context, status int, message string) {
	c.JSON(status, message)
}
