package tradeogre

import (
	"fmt"

	"github.com/tebben/weybot/utils"
)

var (
	endpointAPI    = "https://tradeogre.com/api/v1"
	endpointTicker = fmt.Sprintf("%s/ticker/", endpointAPI)
)

// GetTicker returns the ticker info of a market from tradeogre
func GetTicker(market string) (*Ticker, error) {
	ticker := &Ticker{}
	err := utils.GetJSON(fmt.Sprintf("%s%s", endpointTicker, market), ticker)
	if err != nil {
		return nil, err
	}

	return ticker, nil
}
