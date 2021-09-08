package handler

import (
	"github.com/gin-gonic/gin"
	ModelForGin2 "iceCreamApiWithDI/layers/handler/ModelForGin"
	response2 "iceCreamApiWithDI/layers/handler/response"
)

func (h *Handler) SingUp(c *gin.Context) {
	var input ModelForGin2.CreateUser

	if err := c.BindJSON(&input); err != nil {
		response2.ErrorResponse(c, 400, err.Error())
		return
	}
	err := h.service.CreateUser(input)

	if err != nil {
		response2.ErrorResponse(c, 400, "not unique data")
		return
	}

	response2.OkResponse(c, 200, "success!")
}

func (h *Handler) SingIn(c *gin.Context) {
	var input ModelForGin2.GetUser
	if err := c.BindJSON(&input); err != nil {
		response2.ErrorResponse(c, 400, err.Error())
		return
	}
	token, errLogin := h.service.LoginUser(input)
	if errLogin != nil {
		response2.ErrorResponse(c, 400, errLogin.Error())
		return
	}
	c.JSON(200,gin.H{
		"token":token,
	})

}
