package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

func App(c *gin.Context) {
	pull := c.Param("pull")
	bene := c.Param("benefit")
	resumeKey := c.DefaultQuery("r", "")

	var mode uint8
	mode = 0
	if pull == "1" {
		mode += 1
	}
	if bene == "1" {
		mode += 2
	}

	templates := []string{
		"b-push.html",
		"b-pull.html",
		"b+push.html",
		"b+pull.html",
	}

	c.HTML(http.StatusOK, templates[mode], gin.H{
		"resumeKey": resumeKey,
	})
}

func Home(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	mode := rand.Intn(4)

	urls := []string{
		"/app/0/0",
		"/app/0/1",
		"/app/1/0",
		"/app/1/1",
	}

	c.Redirect(302, urls[mode])
}
