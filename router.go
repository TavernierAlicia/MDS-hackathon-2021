package main

import (
	"github.com/gin-gonic/gin"
	//"fmt"
)

func Router() {

	// define router
	router := gin.Default()

	//to include html
	router.LoadHTMLFiles(
		"templates/index.html", 
		"templates/404.html", 
		"templates/articles.html", 
		"templates/article.html", 
		"templates/housing.html", 
		"templates/house.html", 
		"templates/rights.html", 
		"templates/subject.html",
	)

	//to include js
	router.Static("/js", "./js")
	
	//to include css
	router.Static("/css", "./css")
	
	//to include images
	router.Static("/pictures", "./pictures")

	// *** GET ***

	// ERROR
	router.NoRoute(errorPage)

	// INDEX
	router.GET("/", indexPage)
	router.GET("/index", indexPageRedirect)
	router.GET("/accueil", indexPageRedirect)
	router.GET("/home", indexPageRedirect)

	// LOGEMENT
	router.GET("/housing", housingPage)
	router.GET("/house/:houseid", housePage)

	// MES DROITS
	router.GET("/mesdroits", rightsPage)
	router.GET("/mesdroits/:subject", subjectPage)

	// BLOG
	router.GET("/blog", blogPage)
	router.GET("/blog/:articleid", articlePage)


	// *** POST ***



	// RUN ROUTER
	router.Run("127.0.0.1:3000")
}