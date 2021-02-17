package main

import (
	"io/ioutil"
	"fmt"
	"strings"
	"github.com/gin-gonic/gin"
	"net/http"
	"encoding/json"
)

var zone string
var keyword string

// adding try/catch function on this file //

func (tcf Block) Do() {
	if tcf.Finally != nil {
		defer tcf.Finally()
	}
	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r)
			}
		}()
	}
	tcf.Try()
}

//////////////////////////////////



// ERROR
func errorPage(c *gin.Context) {

	//display error page
	data := gin.H{
		"code": 1,
		"message": "This page is unreacheable",
	}

	c.HTML(404, "404.html", data)
}


// INDEX
func indexPage(c *gin.Context) {

	// request API with a try/catch
	Block{
		Try: func() {
			resp, err := http.Get("https://data.enseignementsup-recherche.gouv.fr/api/records/1.0/search/?dataset=fr_crous_logement_france_entiere&q=&rows=3&facet=zone")
			if err != nil {
				log.Error("Cannot get data from API")
			}
			
			defer resp.Body.Close()
			
			body, err := ioutil.ReadAll(resp.Body)
			
			if err != nil {
				log.Error("Cannot read data from API")
			}
			
			data := Preview{}
			
			err = json.Unmarshal(body, &data)
			if err != nil {
				log.Error("unmarshal error")
			}
			c.HTML(200, "index.html", map[string]interface{}{"data": data.Records})
		},
		Catch: func(e Exception) {
			fmt.Printf("Caught %v\n", e)
			errorPage(c)
		},
	}.Do()
}

func indexPageRedirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/")
}

// LOGEMENT
func housingPage(c *gin.Context) {
	// request API with a try/catch
	Block{
		Try: func() {
				resp, err := http.Get("https://data.enseignementsup-recherche.gouv.fr/api/records/1.0/search/?dataset=fr_crous_logement_france_entiere&q=&rows=10000")
			if err != nil {
				log.Error("Cannot get data from API")
			}
			
			defer resp.Body.Close()
			
			body, err := ioutil.ReadAll(resp.Body)
			
			if err != nil {
				log.Error("Cannot read data from API")
			}
				
			data := Preview{}
				
			err = json.Unmarshal(body, &data)
			if err != nil {
				log.Error("unmarshal error")
			}
			c.HTML(200, "housing.html", map[string]interface{}{"data": data.Records, "zone": zone, "keywords": keyword,})
		},
		Catch: func(e Exception) {
			fmt.Printf("Caught %v\n", e)
			errorPage(c)
		},
	}.Do()
}

// LOGEMENT PER ID
func housePage(c *gin.Context) {

	// Get house id
	id := strings.Replace(c.Request.URL.Path, "/house/", "", -1)

	// Request API with a try/catch to avoid error 500
	Block{
		Try: func() {
			resp, err := http.Get("https://data.enseignementsup-recherche.gouv.fr/api/records/1.0/search/?dataset=fr_crous_logement_france_entiere&q=&rows=1&facet=zone&facet=recordid&refine.recordid="+id)
			if err != nil {
				log.Error("Cannot get data from API")
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Error("Cannot read data from API")
			}
		
			data := FullInfos{}
		
			err = json.Unmarshal(body, &data)
			if err != nil {
				log.Error("unmarshal error")
			}
			c.HTML(200, "house.html", map[string]interface{}{"data": data.Records[0]})
		}, 
		Catch: func(e Exception) {
			fmt.Printf("Caught %v\n", e)
			errorPage(c)
		},
	}.Do()
}


// SEARCH HOME
func searchHome(c *gin.Context) {
	c.Request.ParseForm()
	zone = strings.Join(c.Request.PostForm["zone"], " ")
	keyword = strings.Join(c.Request.PostForm["keyword"], " ")

	var req string

	switch (true) {
	case zone == "tous" && keyword == "":
		housingPage(c)
		return
	case (zone != "tous" && keyword == ""):
		req = "https://data.enseignementsup-recherche.gouv.fr/api/records/1.0/search/?dataset=fr_crous_logement_france_entiere&q=&rows=10000&facet=zone&refine.zone="+zone
	case (zone == "tous" && keyword != ""):
		zone = ""
		req = "https://data.enseignementsup-recherche.gouv.fr/api/records/1.0/search/?dataset=fr_crous_logement_france_entiere&rows=10000&q="+keyword
	case (zone != "tous" && keyword != ""):
		req = "https://data.enseignementsup-recherche.gouv.fr/api/records/1.0/search/?dataset=fr_crous_logement_france_entiere&rows=10000&q="+keyword+"&facet=zone&refine.zone="+zone
	default:
		housingPage(c)
		return
	}

		// make new request
		Block{
			Try: func() {
				fmt.Println(req)
					resp, err := http.Get(req)
				if err != nil {
					log.Error("Cannot get data from API")
				}
				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				
				if err != nil {
					log.Error("Cannot read data from API")
				}
				data := Preview{}
					
				err = json.Unmarshal(body, &data)
				if err != nil {
					log.Error("unmarshal error")
				}
				c.HTML(200, "housing.html", map[string]interface{}{"data": data.Records, "zone": zone, "keywords": keyword,})
			},
			Catch: func(e Exception) {
				fmt.Printf("Caught %v\n", e)
				errorPage(c)
			},
		}.Do()
	
}


// MES DROITS
func rightsPage(c *gin.Context) {
	c.HTML(200, "rights.html", nil)
}

func subjectPage(c *gin.Context) {
	// add subject data 4 templating
	c.HTML(200, "subject.html", nil)
}

// FAQ
func faqPage(c *gin.Context) {
	// add subject data 4 templating
	c.HTML(200, "faq.html", nil)
}

// LEGAL
func legalPage(c *gin.Context) {
	// add subject data 4 templating
	c.HTML(200, "legal.html", nil)
}