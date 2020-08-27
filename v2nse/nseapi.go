package v2nse

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func makeRequest(url string) (*http.Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:79.0) Gecko/20100101 Firefox/79.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")

	return client.Do(req)
}

// GetBoardMeetingsListV2 gets the board meetings list from the NSE Website using API calls
func GetBoardMeetingsListV2(stock string) Meetings {
	resp, err := makeRequest(fmt.Sprintf("https://www.nseindia.com/api/quote-equity?symbol=%s&section=corp_info", stock))
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
func GetStockDataV2(stock string) map[string]string {
	resp, err := makeRequest(fmt.Sprintf("https://www.nseindia.com/api/quote-equity?symbol=%s", stock))
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	defer resp.Body.Close()
	var basicInfo BasicInfo

	err = json.NewDecoder(resp.Body).Decode(&basicInfo)
	finalResult := make(map[string]string)

	finalResult["industry"] = strings.ToLower(basicInfo.Meta.Industry)
	fmt.Println(basicInfo)
	finalResult["high"] = string(basicInfo.Price.WeekHighLow.Max) + "/" + string(basicInfo.Price.WeekHighLow.Min)

	corporate := GetBoardMeetingsListV2(stock)

	finalResult["income"] = corporate.Corp.FinancialResults[0].Income
	finalResult["profit"] = corporate.Corp.FinancialResults[0].ProfitBeforeTax
	finalResult["netprofit"] = corporate.Corp.FinancialResults[0].ProfitAfterTax

	return finalResult
}

// GetTradeInfoV2 gets the trade information related to a stock symbol
func GetTradeInfoV2(stock string) map[string]string {
	resp, err := makeRequest(fmt.Sprintf("https://www.nseindia.com/api/quote-equity?symbol=%s&section=trade_info", stock))
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	defer resp.Body.Close()
	var tradeInfo TradeInfo

	err = json.NewDecoder(resp.Body).Decode(&tradeInfo)
	fmt.Println(tradeInfo)
	finalResult := make(map[string]string)

	finalResult["marketCap"] = tradeInfo.TradeDetails.TradeDetail.TotalMarketCap.String()
	finalResult["volume"] = tradeInfo.TradeDetails.TradeDetail.TotalTradedVolume.String()
	finalResult["value"] = tradeInfo.TradeDetails.TradeDetail.TotalTradedValue.String()

	return finalResult
}
