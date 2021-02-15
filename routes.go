package main

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)


// ERROR
//handle custom error 404
func errorPage(c *gin.Context) {

	//display error page
	data := gin.H{
		"code": 1,
		"message": "This page is unreacheable",
	}
	c.HTML(404, "404.html", data)
}


// INDEX
//handle index page
func indexPage(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

//handle index page redirection
func indexPageRedirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/")
}

// LOGEMENT


// DROITS


// FAQ