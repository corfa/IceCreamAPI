package handler

import (
	"github.com/gin-gonic/gin"
	"iceCreamApiWithDI/handler/ModelForGin"
	"iceCreamApiWithDI/handler/response"
)

func (h *Handler) SingUp(c *gin.Context) {
	var input ModelForGin.CreateUser

	if err := c.BindJSON(&input); err != nil {
		response.ErrorResponse(c, 400, err.Error())
		return
	}
	err := h.service.CreateUser(input)

	if err != nil {
		response.ErrorResponse(c, 400, "not unique data")
		return
	}

	response.OkResponse(c, 200, "success!")
}

func (h *Handler) SingIn(c *gin.Context) {
	var input ModelForGin.GetUser
	if err := c.BindJSON(&input); err != nil {
		response.ErrorResponse(c, 400, err.Error())
		return
	}
	token, errLogin := h.service.LoginUser(input)
	if errLogin != nil {
		response.ErrorResponse(c, 400, errLogin.Error())
		return
	}
	response.OkResponse(c, 200, token)

}
