package tradeogre

// Ticker holds info on volume, high, and low are in the last 24 hours, initialprice is the price from 24 hours ago.
type Ticker struct {
	InitialPrice string `json:"initialprice"`
	Price        string `json:"price"`
	High         string `json:"high"`
	Low          string `json:"low"`
	Volume       string `json:"volume"`
}
