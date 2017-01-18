package main

import (
	"fmt"
	"github.com/bitcoin_ticker/sources"
)

type Display struct {
	BTCUSDSET []float64
	EURUSDSET []float64
	BTCEURSET []float64
	SourceCount int
}

func main() {
	dataChannel := make(chan sources.DataSource)
	fmt.Println("Starting...")
	for _, s := range sources.Sources {
		sources.Run(s, dataChannel)
	}
	sourceCount := len(sources.Sources)
	display := Display{
		BTCUSDSET: make([]float64, sourceCount),
		EURUSDSET: make([]float64, sourceCount),
		BTCEURSET: make([]float64, sourceCount),
		SourceCount: sourceCount,
	}

	for {
		select {
		case dataSource := <-dataChannel:
			go func() {
				display.Print(dataSource)
				
			}()
		}
	}
}

func (d *Display) Print(dataSource sources.DataSource) {
	d.BTCUSDSET[dataSource.GetIndex()] = dataSource.GetBTCUSD()
	btcUsdAvrg, activeSourcesBtcUsd := GetAVRG(d.BTCUSDSET)
	d.EURUSDSET[dataSource.GetIndex()] = dataSource.GetEURUSD()
	eurUsdAvrg, activeSourcesEurUsd := GetAVRG(d.EURUSDSET)
	d.BTCEURSET[dataSource.GetIndex()] = dataSource.GetBTCEUR()
	btcEurAvrg, activeSourcesBtcEur := GetAVRG(d.BTCEURSET)
	fmt.Println("---------------------------------------------------------")
	fmt.Println(fmt.Sprintf("BTC/USD: %f EUR/USD: %f BTC/EUR: %f\n", 
							 btcUsdAvrg,
							 eurUsdAvrg,
							 btcEurAvrg))
	fmt.Println(fmt.Sprintf("Active sources: BTC/USD (%d of %d) EUR/USD (%d of %d) EUR/USD (%d of %d)", 
							 activeSourcesBtcUsd, d.SourceCount,
							 activeSourcesEurUsd, d.SourceCount,
							 activeSourcesBtcEur, d.SourceCount))
	fmt.Println("---------------------------------------------------------\n")
}

func GetAVRG(array []float64) (float64, int) {
	var total float64 = 0
	var counter int = 0
	for _, value := range array {
		total += value
		if value != 0 {
			counter += 1
		}
	}
	return total/float64(counter), counter
}

