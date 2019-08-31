package controllers


import (
	"github.com/gin-gonic/gin"
	"../services"
	"net/http"
	"strconv"
	"../utils"
)

const (
	paramUserId = "userId"
)

func GetUserFromApi(c * gin.Context)  {
	userId, convError:= strconv.Atoi(c.Param(paramUserId))

	if convError != nil {
		apiError := utils.ApiError{
			Message:convError.Error(),
			Status: http.StatusBadRequest,
		}
		c.JSON(apiError.Status, apiError)
		return
	}

	response, err := services.GetUser(userId)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK,response)
}
