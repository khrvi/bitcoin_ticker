package sources

import (
	"strconv"
	"encoding/json"
)

type (
	responseParamsBitkonan struct {
		Last      string `json:"last"`
		High string `json:"high"`
		Low       string `json:"low"`
		Bid       string `json:"bid"`
		Ask       string `json:"ask"`
		Open       string `open:"low"`
		Volume       string `volume:"low"`
	}

	Bitkonan struct {
		Name   string
		BTCUSD float64
		BTCEUR float64
		EURUSD float64
		Index int
		responseParamsBitkonan
	}
)

// {"last":"929.00","high":"929.00","low":"901.00","bid":"878.32","ask":"912.97","open":"930.00","volume":"1.93098257"}
func (b *Bitkonan) GetPath() string {
	return "https://bitkonan.com/api/ticker"
}

func (b *Bitkonan) GetName() string {
	return b.Name
}

func (b *Bitkonan) GetBTCUSD() float64 {
	return b.BTCUSD
}

func (b *Bitkonan) GetBTCEUR() float64 {
	return b.BTCEUR
}

func (b *Bitkonan) GetEURUSD() float64 {
	return b.EURUSD
}

func (b *Bitkonan) GetIndex() int {
	return b.Index
}

func (b *Bitkonan) HandleResponse(body []byte) error {
	err := json.Unmarshal(body, &b.responseParamsBitkonan)
	if err != nil {
		return err
	}
	value, _ := strconv.ParseFloat(b.responseParamsBitkonan.Bid, 64)
	b.BTCUSD = value

	return nil
}

func init() {
	Sources = append(Sources, &Bitkonan{Name: "BitKonan", Index: len(Sources)})
}
