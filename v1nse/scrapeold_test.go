package v1nse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthGetStockResultDetailsV1(t *testing.T) {
	results := GetStockResultDetailsV1("INFY")

	assert.Equal(t, 4, len(results), "Expected Number of values")
}

func TestContentGetStockResultDetailsV1Fail(t *testing.T) {
	expected := []string{"1542500.00", "2032500.00", "490000.00", "400800.00"}

	got := GetStockResultDetailsV1("INFY")
	assert.NotEqual(t, expected, got, "Values of Contents are mismatched shouldnt return the same")
}
