package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"iceCreamApiWithDI/layers/handler/ModelForGin"
	"iceCreamApiWithDI/layers/handler/response"
)

func (h *Handler) addIceCream(c *gin.Context)  {
	var token  ModelForGin.HeaderAuth
	if errToken := c.ShouldBindHeader(&token); errToken != nil {
		c.JSON(400, errors.New("token not found"))
	}

	var input ModelForGin.IceCream

	if errInput := c.BindJSON(&input); errInput != nil {
		response.ErrorResponse(c, 400, errInput.Error())
		return
	}
	errAppendIceCream := h.service.AppendIceCream(input,token)
	if errAppendIceCream != nil {
		response.ErrorResponse(c, 400, errAppendIceCream.Error())
		return
	}

	response.OkResponse(c, 200, "success!")
}
func (h *Handler) DelIceCream(c *gin.Context)  {
	var token  ModelForGin.HeaderAuth
	if errToken := c.ShouldBindHeader(&token); errToken != nil {
		c.JSON(400, errors.New("token not found"))
	}

	var input ModelForGin.IceCream

	if errInput := c.BindJSON(&input); errInput != nil {
		response.ErrorResponse(c, 400, errInput.Error())
		return
	}
	errDeleteIceCream := h.service.DeleteIceCream(input,token)
	if errDeleteIceCream != nil {
		response.ErrorResponse(c, 400, errDeleteIceCream.Error())
		return
	}

	response.OkResponse(c, 200, "success!")
}
