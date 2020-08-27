package v2nse

import "encoding/json"

// Stock is a datatype for Symbol of the stock
type Stock struct {
	Symbol string `json:"symbol"`
}

// Meetings is a custom type to access all the corporate details of a stock through NSE API
type Meetings struct {
	Corp Corporate `json:"corporate"`
}

// Corporate is an embedding for the json strucutre
type Corporate struct {
	BoardMeetings    []BoardMeetings   `json:"boardMeetings"`
	FinancialResults []FinancialResult `json:"financialResults"`
}

// TradeInfo is a wrapper for TradeDetails
type TradeInfo struct {
	TradeDetails TradeDetailsWrapper `json:"marketDeptOrderBook"`
}

// BasicInfo gives the basic details of the stock symbol
type BasicInfo struct {
	Meta  Metadata  `json:"metadata"`
	Price PriceInfo `json:"priceInfo"`
}

// BoardMeetings is a custom type representing the details of board meetings
type BoardMeetings struct {
	BMPurpose   string `json:"bm_purpose"`
	BMDesc      string `json:"bm_desc"`
	Attachment  string `json:"attachment"`
	BMDate      string `json:"bm_date"`
	BMTimestamp string `json:"bm_timestamp"`
}

// Metadata gives the metadata of the stock
type Metadata struct {
	Industry string `json:"industry"`
}

// PriceInfo gives Price Details
type PriceInfo struct {
	WeekHighLow Week52 `json:"weekHighLow"`
}

// Week52 tells the 52 week high/low prices
type Week52 struct {
	Min json.Number `json:"min"`
	Max json.Number `json:"max"`
}

// FinancialResult gives the required detail of each financial result in the json object
type FinancialResult struct {
	Income          string `json:"income"`
	ProfitBeforeTax string `json:"reProLossBefTax"`
	ProfitAfterTax  string `json:"proLossAftTax"`
}

// TradeDetailsWrapper wrapper for TradeDetails struct
type TradeDetailsWrapper struct {
	TradeDetail TradeDetails `json:"tradeInfo"`
}

// TradeDetails gives the traded details information
type TradeDetails struct {
	TotalTradedVolume json.Number `json:"totalTradedVolume"`
	TotalTradedValue  json.Number `json:"totalTradedValue"`
	TotalMarketCap    json.Number `json:"totalMarketCap"`
}
