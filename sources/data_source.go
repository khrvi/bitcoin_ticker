package sources

import (
	"io/ioutil"
	"time"
	"net/http"
)

type DataSource interface {
	GetName() string
	GetBTCUSD() float64
	GetBTCEUR() float64
	GetEURUSD() float64
	GetPath() string
	GetIndex() int
	HandleResponse([]byte) error
}

var Sources = make([]DataSource, 0)

func Run(ds DataSource, dataChannel chan DataSource) error {
	go func(ds DataSource) {
		_ = StartFetching(ds, dataChannel)
		ticker := time.NewTicker(time.Duration(60) * time.Second)
		defer ticker.Stop()
		for _ = range ticker.C {
			var err = StartFetching(ds, dataChannel)
			if err != nil {
				continue
			}
		}
	}(ds)
	return nil
}

func StartFetching(ds DataSource, dataChannel chan DataSource) error {
	path := ds.GetPath()
	resp, err := http.Get(path)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := ds.HandleResponse(body); err != nil {
		return err
	}
	dataChannel <- ds
	return nil
}