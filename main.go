package main

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/hjoshi123/NSEBoardMeetings/v1nse"
	"github.com/hjoshi123/NSEBoardMeetings/v2nse"
)

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

	v1 := router.Group("/api/v1")
	{
		v1.GET("/boardMeetingsList", getBoardMeetingsListV1)
		v1.GET("/stockDetails", getStockDetailsV1)
	}

	v2 := router.Group("/api/v2")
	{
		v2.GET("/boardMeetingsList", getBoardMeetingsListV2)
		v2.GET("/stockDetails", getStockDetailsV2)
	}

	router.Run(":" + port)
}

// getBoardMeetingsListV1 gets the board meetings list from the NSE www1 (old) website due to be stopped on Aug 2020
func getBoardMeetingsListV1(c *gin.Context) {
	stock := c.Query("symbol")

	meetingDate, meetingPurpose := v1nse.GetBoardMeetingsV1(stock)
	c.JSON(http.StatusOK, gin.H{
		"dates":   meetingDate,
		"purpose": meetingPurpose,
	})
}

// getBoardMeetingsListV2 gets the board meetings list from the NSE Website using API calls
func getBoardMeetingsListV2(c *gin.Context) {
	stock := c.Query("symbol")

	corporate := v2nse.GetBoardMeetingsListV2(stock)

	c.JSON(http.StatusOK, gin.H{
		"result": corporate,
	})
}

func getStockDetailsV1(c *gin.Context) {
	stock := c.Query("symbol")

	basics := v1nse.GetStockBasicDetailsV1(stock)
	findata := v1nse.GetStockResultDetailsV1(stock)

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

func getStockDetailsV2(c *gin.Context) {
}
