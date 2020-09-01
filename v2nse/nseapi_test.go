package v2nse

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBoardMeetingsListInvalidInput(t *testing.T) {
	result, err := GetBoardMeetingsListV2("invalid")

	assert.Equal(t, err, nil, "No Errors")
	assert.Equal(t, result, Meetings{}, "Expected result to have empty interface since invalid input")
}

func TestGetStockDataV2NoFinancialResults(t *testing.T) {
	result, err := GetStockDataV2("NIFTYBEES")

	assert.Equal(t, err, nil, "No Error since its a valid input")

	expected := []string{"N/A", "N/A", "N/A"}
	actual := []string{result["income"], result["profit"], result["netprofit"]}

	assert.Equal(t, actual, expected, "Expected Length is 2 since its an ETF")
}

func TestGetTradeInfoV2NoInput(t *testing.T) {
	_, err := GetTradeInfoV2("")

	assert.Equal(t, err, errors.New("Empty Input"), "Error exists as empty input isnt allowed")
}
