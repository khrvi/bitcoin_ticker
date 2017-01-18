package sources

import (
	"encoding/json"
)

type (
	responseParamsСoindesk struct {
		Time      interface{} `json:"time"`
		Disclaimer string `json:"disclaimer"`
		Bpi       map[string]interface{} `json:"bpi"`
	}

	Сoindesk struct {
		Name   string
		BTCUSD float64
		BTCEUR float64
		EURUSD float64
		Index int
		responseParamsСoindesk
	}
)

// {"time":{"updated":"Jan 18, 2017 19:34:00 UTC","updatedISO":"2017-01-18T19:34:00+00:00","updateduk":"Jan 18, 2017 at 19:34 GMT"},
//  "disclaimer":"This data was produced from the CoinDesk Bitcoin Price Index (USD). Non-USD currency data converted using hourly conversion rate from openexchangerates.org",
//  "bpi":{"USD":{"code":"USD","symbol":"&#36;","rate":"867.5175","description":"United States Dollar","rate_float":867.5175},
//         "GBP":{"code":"GBP","symbol":"&pound;","rate":"705.6717","description":"British Pound Sterling","rate_float":705.6717},
//         "EUR":{"code":"EUR","symbol":"&euro;","rate":"813.3193","description":"Euro","rate_float":813.3193}}}
func (c *Сoindesk) GetPath() string {
	return "http://api.coindesk.com/v1/bpi/currentprice.json"
}

func (c *Сoindesk) GetName() string {
	return c.Name
}

func (c *Сoindesk) GetBTCUSD() float64 {
	return c.BTCUSD
}

func (c *Сoindesk) GetBTCEUR() float64 {
	return c.BTCEUR
}

func (c *Сoindesk) GetEURUSD() float64 {
	return c.EURUSD
}

func (c *Сoindesk) GetIndex() int {
	return c.Index
}

func (c *Сoindesk) HandleResponse(body []byte) error {
	err := json.Unmarshal(body, &c.responseParamsСoindesk)
	if err != nil {
		return err
	}
	c.BTCUSD = c.responseParamsСoindesk.Bpi["USD"].(map[string]interface{})["rate_float"].(float64)
	c.BTCEUR = c.responseParamsСoindesk.Bpi["EUR"].(map[string]interface{})["rate_float"].(float64)

	return nil
}

func init() {
	Sources = append(Sources, &Сoindesk{Name: "СoinDesk", Index: len(Sources)})
}
