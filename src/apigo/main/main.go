package main

import (
	"github.com/gin-gonic/gin"
	"../controllers"
	"log"
)

const (
	port = ":8082"
)

var (
	router = gin.Default()
)

func main()  {

	router.GET("/results/:userId", controllers.GetResultFromApi)
	router.GET("/results/wg/:userId", controllers.GetResultFromApiW)
	router.GET("/results/ch/:userId", controllers.GetResultFromApiC)

	log.Fatal(router.Run(port))

}
