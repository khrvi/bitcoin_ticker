package sources

import (
	"fmt"
	"strconv"
	"encoding/json"
)

const endpoint = "https://api.bitfinex.com/v1"

type (
	responseParams struct {
		Mid       string `json:"mid"`
		Bid       string `json:"bid"`
		Ask       string `json:"ask"`
		LastPrice string `json:"last_price"`
		Low       string `json:"low"`
		High      string `json:"high"`
		Volume    string `json:"volume"`
		Timestamp string `json:"timestamp"`
	}

	Bitfinex struct {
		Name   string
		BTCUSD float64
		BTCEUR float64
		EURUSD float64
		Index int
		responseParams
	}
)

// {"mid":"895.46","bid":"895.11","ask":"895.81","last_price":"895.94","low":"875.0","high":"913.85","volume":"21104.45536449","timestamp":"1484743822.597969449"}
func (b *Bitfinex) GetPath() string {
	return fmt.Sprintf("%s/pubticker/btcusd", endpoint)
}

func (b *Bitfinex) GetName() string {
	return b.Name
}

func (b *Bitfinex) GetBTCUSD() float64 {
	return b.BTCUSD
}

func (b *Bitfinex) GetBTCEUR() float64 {
	return b.BTCEUR
}

func (b *Bitfinex) GetEURUSD() float64 {
	return b.EURUSD
}

func (b *Bitfinex) GetIndex() int {
	return b.Index
}

func (d *Bitfinex) HandleResponse(body []byte) error {
	err := json.Unmarshal(body, &d.responseParams)
	if err != nil {
		return err
	}

	value, _ := strconv.ParseFloat(d.responseParams.Mid, 64)
	d.BTCUSD = value

	return nil
}

func init() {
	Sources = append(Sources, &Bitfinex{Name: "Bitfinex", Index: len(Sources)})
}
