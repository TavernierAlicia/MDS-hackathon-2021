package main

import (
	"github.com/gin-gonic/gin"
	//"fmt"
)

func Router() {

	// define router
	router := gin.Default()

	//to include html
	router.LoadHTMLFiles("templates/index.html", "templates/404.html")

	//to include js
	router.Static("/js", "./js")
	
	//to include css
	router.Static("/css", "./css")
	
	//to include images
	router.Static("/pictures", "./pictures")

	// GET

	// ERROR
	router.NoRoute(errorPage)

	// INDEX
	router.GET("/", indexPage)
	router.GET("/index", indexPageRedirect)
	router.GET("/accueil", indexPageRedirect)
	router.GET("/home", indexPageRedirect)




	// POST

	// RUN ROUTER
	router.Run("127.0.0.1:3000")
}