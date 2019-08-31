package controllers


import (
	"github.com/gin-gonic/gin"
	"../services"
	"net/http"
	"strconv"
	"../utils"
	"sync"
	"../patterns"
)

func GetResultFromApi(c * gin.Context)  {

	userId, convError:= strconv.Atoi(c.Param(paramUserId))

	if convError != nil {
		apiError := utils.ApiError{
			Message:convError.Error(),
			Status: http.StatusBadRequest,
		}
		c.JSON(apiError.Status, apiError)
		return
	}

	response := services.GetResult(userId)

	if response.ApiError != nil {
		c.JSON(response.ApiError.Status, response.ApiError)
		return
	}

	c.JSON(http.StatusOK,response)
}

func GetResultFromApiW(c * gin.Context)  {

	userId, convError:= strconv.Atoi(c.Param(paramUserId))

	if convError != nil {
		apiError := utils.ApiError{
			Message:convError.Error(),
			Status: http.StatusBadRequest,
		}
		c.JSON(apiError.Status, apiError)
		return
	}

	var wg sync.WaitGroup

	response := services.GetResultW(userId, wg)

	if response.ApiError != nil {
		c.JSON(response.ApiError.Status, response.ApiError)
		return
	}

	c.JSON(http.StatusOK,response)
}

func GetResultFromApiC(c * gin.Context)  {

	userId, convError:= strconv.Atoi(c.Param(paramUserId))
	cb := patterns.CircuitBreaker{}

	patterns.IniCircuit(cb, 4)

	if convError != nil {
		apiError := utils.ApiError{
			Message:convError.Error(),
			Status: http.StatusBadRequest,
		}
		c.JSON(apiError.Status, apiError)
		return
	}

	if patterns.ConnectionEnabled(cb){
		response := services.GetResultC(userId)

		if response.ApiError != nil {
			patterns.InformError(cb)
			c.JSON(response.ApiError.Status, response.ApiError)
			return
		}

		c.JSON(http.StatusOK,response)
	}



	c.JSON(http.StatusInternalServerError, utils.ApiError{
		Message:"No anda nada, esta todo roto",
		Status: 500,
	})



}

