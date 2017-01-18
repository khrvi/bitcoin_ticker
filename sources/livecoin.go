package sources

import (
	"encoding/json"
)

type (
	Livecoin struct {
		Name   string
		BTCUSD float64
		BTCEUR float64
		EURUSD float64
		Index int
	}
)

// [{"cur":"USD","symbol":"USD/RUR","last":55.10000000,"high":61.98000000,"low":55.10000000,"volume":702.16598814,"vwap":61.32264048,"max_bid":58.30000000,"min_ask":55.10000000,"best_bid":55.55500000,"best_ask":59.90000000},
//  {"cur":"EUR","symbol":"EUR/USD","last":1.04000000,"high":1.04400000,"low":1.04000000,"volume":6.00000000,"vwap":1.04361311,"max_bid":1.04400000,"min_ask":0,"best_bid":1.04200000,"best_ask":1.10000000},
//  ...
func (l *Livecoin) GetPath() string {
	return "https://api.livecoin.net/exchange/ticker"
}

func (l *Livecoin) GetName() string {
	return l.Name
}

func (l *Livecoin) GetBTCUSD() float64 {
	return l.BTCUSD
}

func (l *Livecoin) GetBTCEUR() float64 {
	return l.BTCEUR
}

func (l *Livecoin) GetEURUSD() float64 {
	return l.EURUSD
}

func (l *Livecoin) GetIndex() int {
	return l.Index
}

func (l *Livecoin) HandleResponse(body []byte) error {
	var resources []interface{}
	err := json.Unmarshal(body, &resources)
	if err != nil {
		return err
	}
	for _, resource := range resources {
		if resource.(map[string]interface{})["symbol"] == "BTC/USD" {
			l.BTCUSD = resource.(map[string]interface{})["vwap"].(float64)
		}
		if resource.(map[string]interface{})["symbol"] == "BTC/EUR" {
			l.BTCEUR = resource.(map[string]interface{})["vwap"].(float64)
		}
		if resource.(map[string]interface{})["symbol"] == "EUR/USD" {
			l.EURUSD = resource.(map[string]interface{})["vwap"].(float64)
		}	
	}
	

	return nil
}

func init() {
	Sources = append(Sources, &Livecoin{Name: "Livecoin", Index: len(Sources)})
}
