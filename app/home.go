package app

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

var g *gin.Engine

func Home(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())

	var pull, sso bool
	var url string

	mode := rand.Intn(4)

	if mode == 0 {
		pull = false
		sso = false
		url = "https://marvelapp.com/3f188jg"
	}
	if mode == 1 {
		pull = true
		sso = false

		url = "https://marvelapp.com/f05ei9"
	}
	if mode == 2 {
		pull = false
		sso = true
		url = "https://marvelapp.com/3e7edg6"
	}
	if mode == 3 {
		pull = true
		sso = true

		url = "https://marvelapp.com/f05c23"
	}

	c.HTML(http.StatusOK, "layout.html", gin.H{
		"pull": pull,
		"sso":  sso,
		"url":  url,
	})
}

func NotFound(c *gin.Context) {
	c.JSON(404, gin.H{"status": "not found"})
}
