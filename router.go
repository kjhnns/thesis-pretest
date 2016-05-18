package main

import (
	"fmt"
)

func initRouting() {

	fmt.Println("Routing")
	g.Static("/public", "./public")
	g.LoadHTMLGlob("templates/*")
	g.StaticFile("/favicon.ico", "./public/images/favicon.png")

	g.GET("/chat", Chat)
	g.GET("/coupon", Coupon)

}
