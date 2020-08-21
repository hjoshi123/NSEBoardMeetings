package v2nse

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// GetBoardMeetingsListV2 gets the board meetings list from the NSE Website using API calls
func GetBoardMeetingsListV2(stock string) Meetings {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://www.nseindia.com/api/quote-equity?symbol=%s&section=corp_info", stock), nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:79.0) Gecko/20100101 Firefox/79.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	defer resp.Body.Close()

	var corporate Meetings

	err = json.NewDecoder(resp.Body).Decode(&corporate)

	if err != nil {
		fmt.Println(err)
	}

	return corporate
}

// GetStockDataV2 gets the board meetings list from the NSE Website using API calls
func GetStockDataV2(stock string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://www.nseindia.com/api/quote-equity?symbol=%s", stock), nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:79.0) Gecko/20100101 Firefox/79.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	defer resp.Body.Close()
	var basicInfo BasicInfo

	err = json.NewDecoder(resp.Body).Decode(&basicInfo)
	finalResult := make(map[string]string)

	finalResult["industry"] = basicInfo.Meta.Industry
	finalResult["high"] = basicInfo.Price.HighLow.Max + "/" + basicInfo.Price.HighLow.Min
}
