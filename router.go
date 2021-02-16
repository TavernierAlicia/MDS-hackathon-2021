package main

import (
	"github.com/gin-gonic/gin"
	//"fmt"
)

func Router() {

	// define router
	router := gin.Default()

	//to include html
	router.LoadHTMLGlob("templates/*")

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
	router.GET("/rights", rightsPage)
	router.GET("/rights/:subject", subjectPage)

	// FAQ
	router.GET("/faq", faqPage)

	// LEGAL
	router.GET("/legal", legalPage)

	// *** POST ***



	// RUN ROUTER
	router.Run("127.0.0.1:3000")
}