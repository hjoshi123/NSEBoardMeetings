package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
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
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"POST", "GET"},
		AllowHeaders:  []string{"Origin, Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "Welcome to API",
		})
	})

	router.POST("/api/boardMeetingsList", getBoardMeetingsList)
	router.POST("/api/stockDetails", getStockDetails)
	router.Run(":" + port)
}

func getBoardMeetingsList(c *gin.Context) {
	var stock Stock
	c.BindJSON(&stock)

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

	col.OnHTML(`body`, func(e *colly.HTMLElement) {
		var meetingDate []string
		var meetingPurpose []string
		e.ForEach("table tbody tr td table tbody tr td table", func(_ int, el *colly.HTMLElement) {
			el.ForEach("td:nth-of-type(1)", func(_ int, el *colly.HTMLElement) {
				meetingDate = append(meetingDate, el.Text)
			})
			el.ForEach("td:nth-of-type(2)", func(_ int, el *colly.HTMLElement) {
				meetingPurpose = append(meetingPurpose, el.Text)
			})
		})

		if len(meetingDate) >= 11 && len(meetingPurpose) >= 11 {
			// fetch the only top 10 meeting dates
			c.JSON(http.StatusOK, gin.H{
				"dates":   meetingDate[1:11],
				"purpose": meetingPurpose[1:11],
			})
		} else if len(meetingDate) <= 10 && len(meetingPurpose) <= 10 {
			c.JSON(http.StatusOK, gin.H{
				"dates":   meetingDate,
				"purpose": meetingPurpose,
			})
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"dates":   make([]string, 0),
				"purpose": make([]string, 0),
			})
		}
	})

	col.Visit(fmt.Sprintf("https://www1.nseindia.com/marketinfo/companyTracker/boardMeeting.jsp?symbol=" + stock.Symbol))
}

func getStockDetails(c *gin.Context) {
	var stock Stock
	c.BindJSON(&stock)

	basics := getBasicStockDetails(stock.Symbol)
	findata := getStockResultDetails(stock.Symbol)

	var result []string
	for _, value := range basics {
		str := strings.TrimSpace(strings.Split(value, ":")[1])

		result = append(result, str)
	}

	for _, value := range findata {
		result = append(result, value)
	}

	finalResult := make(map[string]string)
	j := 0

	if len(result) == 3 {
		finalResult["industry"] = "N/A"
		finalResult["high"] = "N/A"
		j = 0
	} else if len(result) == 4 {
		finalResult["industry"] = "N/A"
		finalResult["high"] = result[0]
		j = 1
	} else {
		finalResult["industry"] = strings.ToLower(result[0])
		finalResult["high"] = result[1]
		finalResult["expenditure"] = result[2]
		finalResult["profit"] = result[3]
		finalResult["netprofit"] = result[4]
		j = 4
	}

	if j <= 1 {
		finalResult["expenditure"] = result[j]
		finalResult["profit"] = result[j+1]
		finalResult["netprofit"] = result[j+2]
	}

	c.JSON(http.StatusOK, gin.H{
		"result": finalResult,
	})
}

func getBasicStockDetails(stock string) []string {
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
	var result []string

	col.OnHTML("body", func(e *colly.HTMLElement) {
		e.ForEach("table tbody tr", func(_ int, e *colly.HTMLElement) {
			e.ForEach("td", func(_ int, el *colly.HTMLElement) {
				if strings.Contains(el.Text, "52") || strings.Contains(el.Text, "Industry") {
					result = append(result, el.Text)
				}
			})
		})
	})

	col.Visit(fmt.Sprintf("https://www1.nseindia.com/marketinfo/companyTracker/compInfo.jsp?symbol=%s&series=EQ", stock))
	return result
}

func getStockResultDetails(stock string) []string {
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

	var result []string
	col.OnHTML("body", func(e *colly.HTMLElement) {
		e.ForEach("table tbody tr td table:nth-of-type(3) tbody tr", func(_ int, el *colly.HTMLElement) {
			el.ForEach("td:nth-of-type(2)", func(_ int, el *colly.HTMLElement) {
				result = append(result, el.Text)
			})
		})
	})

	col.Visit(fmt.Sprintf("https://www1.nseindia.com/marketinfo/companyTracker/resultsCompare.jsp?symbol=%s", stock))
	return result[2:5]
}
