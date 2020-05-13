package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	_ "github.com/heroku/x/hmetrics/onload"
)

// Stock is a datatype for Symbol of the stock
type Stock struct {
	Symbol string `json:"symbol"`
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.POST("/getBoardMeetings", getBoardMeetings)

	router.Run(":" + port)
}

func getBoardMeetings(c *gin.Context) {
	var stock Stock
	c.BindJSON(&stock)

	log.Print(stock.Symbol)
	col := colly.NewCollector()

	if runtime.GOOS == "darwin" {
		col.UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:76.0) Gecko/20100101 Firefox/76.0"
	} else if runtime.GOOS == "linux" {
		col.UserAgent = "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:76.0) Gecko/20100101 Firefox/76.0"
		col.UserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36"
		col.UserAgent = "Mozilla/5.0 (Linux x86_64; rv:76.0) Gecko/20100101 Firefox/76.0"
	}

	// Before making a request print "Visiting ..."
	col.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
		log.Println("visiting", r.URL.String())
	})

	col.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	var boardMeetingDetails map[string]interface{}

	col.OnHTML(`body`, func(e *colly.HTMLElement) {
		fmt.Println("Hello")
		fmt.Printf("Link found: %q\n", e.Text)

		pattern := regexp.MustCompile("(\\\"(.*?)\\\"|(\\w+))(\\s*:\\s*(\\\".*?\\\"|.))")
		jsonString := pattern.ReplaceAllString(e.Text, "\"$2$3\"$4")

		log.Println(jsonString)

		if err := json.Unmarshal([]byte(jsonString), &boardMeetingDetails); err != nil {
			panic(err)
		}

		if boardMeetingDetails["results"] == 0 {
			c.JSON(http.StatusOK, gin.H{
				"result": make([]string, 0),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"result": boardMeetingDetails["rows"],
			})
		}
	})

	col.Visit("https://www1.nseindia.com/corporates/corpInfo/equities/getBoardMeetings.jsp?Symbol=" + stock.Symbol + "&Industry=&Purpose=&period=Latest%20Announced&symbol=" + stock.Symbol + "&industry=&purpose=")
}
