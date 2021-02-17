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
var allnil int

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
			resp, err := http.Get("https://data.enseignementsup-recherche.gouv.fr/api/records/1.0/search/?dataset=fr_crous_logement_france_entiere&q=&rows=4&facet=zone")
			if err != nil {
				log.Error("Cannot get data from API")
			}
			
			defer resp.Body.Close()
			
			body, err := ioutil.ReadAll(resp.Body)
			
			if err != nil {
				log.Error("Cannot read data from API")
			}
			
			houses := Preview{}
			
			err = json.Unmarshal(body, &houses)
			if err != nil {
				log.Error("unmarshal error")
			}

			// request database
			db, err := RunDb()
			if err != nil {
				log.Error("Cannot connect to database")
			}
			articles := []*ArticleInfos{}
			err = db.Select(&articles, "SELECT article.id, article.title, article.main_picture, CONCAT(LEFT(article_content.text, 100), '...') AS text, 0.0 AS relevance FROM article JOIN article_content ON article.id = article_content.article_id GROUP BY article.id LIMIT 4")
			
			fmt.Println(articles)
			c.HTML(200, "index.html", map[string]interface{}{"houses": houses.Records, "articles": articles})
		},
		Catch: func(e Exception) {
			log.Error("Caught %v\n "+fmt.Sprintf("%v", e))
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
				resp, err := http.Get("https://data.enseignementsup-recherche.gouv.fr/api/records/1.0/search/?dataset=fr_crous_logement_france_entiere&q=&rows=1000")
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
			log.Error("Caught %v\n "+fmt.Sprintf("%v", e))
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
			log.Error("Caught %v\n "+fmt.Sprintf("%v", e))
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

	// switch case for request type
	switch (true) {
	case zone == "tous" && keyword == "":
		housingPage(c)
		return
	case (zone != "tous" && keyword == ""):
		req = "https://data.enseignementsup-recherche.gouv.fr/api/records/1.0/search/?dataset=fr_crous_logement_france_entiere&q=&rows=1000&facet=zone&refine.zone="+zone
	case (zone == "tous" && keyword != ""):
		zone = ""
		req = "https://data.enseignementsup-recherche.gouv.fr/api/records/1.0/search/?dataset=fr_crous_logement_france_entiere&rows=1000&q="+keyword
	case (zone != "tous" && keyword != ""):
		req = "https://data.enseignementsup-recherche.gouv.fr/api/records/1.0/search/?dataset=fr_crous_logement_france_entiere&rows=1000&q="+keyword+"&facet=zone&refine.zone="+zone
	default:
		housingPage(c)
		return
	}

		// make new request
		Block{
			Try: func() {
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
				log.Error("Caught %v\n "+fmt.Sprintf("%v", e))
				errorPage(c)
			},
		}.Do()
	
}


// MES DROITS
func rightsPage(c *gin.Context) {
	// request database
	db, err := RunDb()
	if err != nil {
		log.Error("Cannot connect to database")
	}
	data := []*ArticleInfos{}
	db.Select(&data, "SELECT article.id, article.title, article.main_picture, CONCAT(LEFT(article_content.text, 100), '...') AS text, 0.0 AS relevance FROM article JOIN article_content ON article.id = article_content.article_id GROUP BY article.id LIMIT 4")
	c.HTML(200, "rights.html", map[string]interface{}{"data": data, "keywords": keyword,})
}

// Search Article
func searchRights(c *gin.Context) {
	// manage keywords
	c.Request.ParseForm()
	keywords := strings.Join(c.Request.PostForm["keyword"], " ")
	keywords = strings.Replace(keywords, "'", "", -1)

	if keywords == "" {
		rightsPage(c)
	}

	fmt.Println(keywords)
	db, err := RunDb()
	if err != nil {
		log.Error("Cannot connect to database")
	}

	data := []*ArticleInfos{}
	db.Select(&data, "SELECT article_content.article_id AS id, article.title, article.main_picture, CONCAT(LEFT(article_content.text, 100), '...') AS text, SUM(MATCH (article_content.low_title) AGAINST('"+keywords+"*') + MATCH (article_content.text) AGAINST('"+keywords+"*') + MATCH (article.title) AGAINST ('"+keywords+"*')) as relevance FROM article_content JOIN article ON article_content.id = article.id WHERE MATCH (article_content.low_title) AGAINST('"+keywords+"*' IN BOOLEAN MODE) OR MATCH (article_content.text) AGAINST('"+keywords+"*' IN BOOLEAN MODE) OR MATCH (article.title) AGAINST ('"+keywords+"*' IN BOOLEAN MODE) GROUP BY article_id ORDER BY relevance")
	c.HTML(200, "rights.html", map[string]interface{}{"data": data, "keywords": keywords,})
	
	// keywords precious request
	// SELECT article_content.article_id AS id, article.title, article.main_picture, CONCAT(LEFT(article_content.text, 100), '...') AS text, SUM(MATCH (article_content.low_title) AGAINST('"+keywords+"*') + MATCH (article_content.text) AGAINST('"+keywords+"*') + MATCH (article.title) AGAINST ('"+keywords+"*')) as relevance FROM article_content JOIN article ON article_content.id = article.id WHERE MATCH (article_content.low_title) AGAINST('"+keywords+"*' IN BOOLEAN MODE) OR MATCH (article_content.text) AGAINST('"+keywords+"*' IN BOOLEAN MODE) OR MATCH (article.title) AGAINST ('"+keywords+"*' IN BOOLEAN MODE) GROUP BY article_id ORDER BY relevance 

}


// ARTICLE
func subjectPage(c *gin.Context) {

	id := strings.Replace(c.Request.URL.Path, "/rights/", "", -1)

	db, err := RunDb()
	if err != nil {
		log.Error("Cannot connect to database")
	}
	
	infos := []*ArticleInfos{}
	db.Select(&infos, "SELECT id, title, main_picture, '' AS text, 0.0 AS relevance FROM article WHERE id = ?", id)


	content := []*ArticleContent{}
	db.Select(&content, "SELECT id, IFNULL(low_title, '') AS low_title, IFNULL(text, '') AS text, IFNULL(picture, '') AS picture FROM article_content WHERE article_id = ? ORDER BY id ASC", id)



	c.HTML(200, "subject.html", map[string]interface{}{"infos": infos, "content": content})
}


// FAQ
func faqPage(c *gin.Context) {
	db, err := RunDb()
	if err != nil {
		errorPage(c)
	} else {
		data := []*FAQ{}
		db.Select(&data, "SELECT * FROM faq")
		c.HTML(200, "faq.html", map[string]interface{}{"data": data})
	}
}

// LEGAL
func legalPage(c *gin.Context) {
	c.HTML(200, "legal.html", nil)
}


// SEARCHALL
func searchAll(c *gin.Context) {
	c.Request.ParseForm()
	keywords := strings.Join(c.Request.PostForm["keyword"], " ")
	keywords = strings.Replace(keywords, "'", "", -1)

	if keywords == "" {
		indexPage(c)
	}

	Block{
		Try: func() {
			resp, err := http.Get("https://data.enseignementsup-recherche.gouv.fr/api/records/1.0/search/?dataset=fr_crous_logement_france_entiere&rows=1000&q="+keywords)
			if err != nil {
				log.Error("Cannot get data from API")
			}
			defer resp.Body.Close()
			
			body, err := ioutil.ReadAll(resp.Body)
			
			if err != nil {
				log.Error("Cannot read data from API")
			}
			
			houses := Preview{}
			
			err = json.Unmarshal(body, &houses)
			if err != nil {
				log.Error("unmarshal error")
			}

			// request database
			db, err := RunDb()
			if err != nil {
				log.Error("Cannot connect to database")
			}
			articles := []*ArticleInfos{}
			db.Select(&articles, "SELECT article_content.article_id AS id, article.title, article.main_picture, CONCAT(LEFT(article_content.text, 100), '...') AS text, SUM(MATCH (article_content.low_title) AGAINST('"+keywords+"*') + MATCH (article_content.text) AGAINST('"+keywords+"*') + MATCH (article.title) AGAINST ('"+keywords+"*')) as relevance FROM article_content JOIN article ON article_content.id = article.id WHERE MATCH (article_content.low_title) AGAINST('"+keywords+"*' IN BOOLEAN MODE) OR MATCH (article_content.text) AGAINST('"+keywords+"*' IN BOOLEAN MODE) OR MATCH (article.title) AGAINST ('"+keywords+"*' IN BOOLEAN MODE) GROUP BY article_id ORDER BY relevance ")
			
			faq := []*FAQ{}
			db.Select(&faq, "SELECT id, question, answer, SUM(MATCH (question) AGAINST('"+keywords+"*') + MATCH (answer) AGAINST('"+keywords+"*')) as relevance FROM faq WHERE MATCH (question) AGAINST('"+keywords+"*' IN BOOLEAN MODE) OR MATCH (answer) AGAINST('"+keywords+"*' IN BOOLEAN MODE) GROUP BY id ORDER BY relevance ")

			allnil = func() int { if len(houses.Records) == 0 && len(articles) == 0 && len(faq) == 0  {return 1 } else {return 0} }()
			
			c.HTML(200, "searchResults.html", map[string]interface{}{"houses": houses.Records, "articles": articles, "faq": faq, "keywords": keywords, "allnil": allnil})
		},
		Catch: func(e Exception) {
			log.Error("Caught %v\n "+fmt.Sprintf("%v", e))
			errorPage(c)
		},
	}.Do()
}
