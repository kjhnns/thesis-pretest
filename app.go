package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Chat(c *gin.Context) {
	c.HTML(http.StatusOK, "chat.html", gin.H{})
}

func Coupon(c *gin.Context) {
	c.HTML(http.StatusOK, "coupon.html", gin.H{})
}
