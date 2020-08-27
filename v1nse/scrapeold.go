package v1nse

import (
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/gocolly/colly"
)

func getCollyInstance() *colly.Collector {
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

	return col
}

// GetStockResultDetailsV1 gets the stocks financial results from the nse's old website
func GetStockResultDetailsV1(stock string) []string {
	col := getCollyInstance()

	var result []string
	col.OnHTML("body", func(e *colly.HTMLElement) {
		e.ForEach("table tbody tr td table:nth-of-type(3) tbody tr", func(_ int, el *colly.HTMLElement) {
			el.ForEach("td:nth-of-type(2)", func(_ int, el *colly.HTMLElement) {
				result = append(result, el.Text)
			})
		})
	})

	col.Visit(fmt.Sprintf("https://www1.nseindia.com/marketinfo/companyTracker/resultsCompare.jsp?symbol=%s", stock))
	return result[1:5]
}

// GetStockBasicDetailsV1 gets the basic stock details like the industry it belongs to etc
func GetStockBasicDetailsV1(stock string) []string {
	col := getCollyInstance()

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

// GetBoardMeetingsV1 takes in stock and returns the list of board meetings from the NSE website
func GetBoardMeetingsV1(stock string) ([]string, []string) {
	col := getCollyInstance()

	var meetingDate []string
	var meetingPurpose []string
	col.OnHTML(`body`, func(e *colly.HTMLElement) {

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
			meetingDate = meetingDate[1:11]
			meetingPurpose = meetingPurpose[1:11]
		} else if len(meetingDate) <= 10 && len(meetingPurpose) <= 10 {
		} else {
			meetingPurpose = make([]string, 0)
			meetingDate = make([]string, 0)
		}
	})

	col.Visit(fmt.Sprintf("https://www1.nseindia.com/marketinfo/companyTracker/boardMeeting.jsp?symbol=" + stock))
	return meetingDate, meetingPurpose
}
