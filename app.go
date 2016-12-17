package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
	"fmt"
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

func Redirect(c *gin.Context) {
	pull := c.Param("pull")
	popular := c.Param("popular")
	disclose := c.Param("disclose")
	url := fmt.Sprintf("http://umfragen.ise.tu-darmstadt.de/sosci/privacyresearch/?password=test&pull=%d&popular=%d&disclose=%d", pull, popular, disclose)
	c.Redirect(http.StatusMovedPermanently, url)
}

func Home(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	mode := rand.Intn(6)

	urls := []string{
		"https://marvelapp.com/17475hc/screen/17622836",
		"https://marvelapp.com/17477hg/screen/17622959",
		"https://marvelapp.com/4f55e1j/screen/17622902",
		"https://marvelapp.com/33a2g74/screen/17622421",
		"https://marvelapp.com/33a3bgh/screen/17622786",
		"https://marvelapp.com/33a3a8e/screen/17622754",
	}

	// urls := []string{
	// 	"/app/0/0",
	// 	"/app/0/1",
	// 	"/app/1/0",
	// 	"/app/1/1",
	// }

	c.Redirect(302, urls[mode])
}
