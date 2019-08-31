package controllers

import (
"github.com/gin-gonic/gin"
"../services"
"net/http"
)

const (
	paramSiteId = "siteId"
)

func GetSiteFromApi(c * gin.Context)  {
	siteId := c.Param(paramSiteId)

	response, err := services.GetSite(siteId)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK,response)
}