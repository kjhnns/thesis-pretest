package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

var g *gin.Engine

func bootstrap() {

	fmt.Println("bootstrapping")

	// Init Gin-Gonic
	g = gin.Default()
	initRouting()

	port := "3000"
	//if PORT passed as env-variable, use the port
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	// Listen and serve on 0.0.0.0:3000
	g.Run(":" + port)
}
