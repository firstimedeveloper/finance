package stock

import (
	"fmt"
	"strings"
	"time"
)

type StockList []Stock

type Stock struct {
	Quote struct {
		Symbol                string  `json:"symbol"`
		CompanyName           string  `json:"companyName"`
		PrimaryExchange       string  `json:"primaryExchange"`
		Sector                string  `json:"sector"`
		CalculationPrice      string  `json:"calculationPrice"`
		Open                  float64 `json:"open"`
		OpenTime              int64   `json:"openTime"`
		Close                 float64 `json:"close"`
		CloseTime             int64   `json:"closeTime"`
		High                  float64 `json:"high"`
		Low                   float64 `json:"low"`
		LatestPrice           float64 `json:"latestPrice"`
		LatestSource          string  `json:"latestSource"`
		LatestTime            string  `json:"latestTime"`
		LatestUpdate          int64   `json:"latestUpdate"`
		LatestVolume          int     `json:"latestVolume"`
		IexRealtimePrice      float64 `json:"iexRealtimePrice"`
		IexRealtimeSize       int     `json:"iexRealtimeSize"`
		IexLastUpdated        int64   `json:"iexLastUpdated"`
		DelayedPrice          float64 `json:"delayedPrice"`
		DelayedPriceTime      int64   `json:"delayedPriceTime"`
		ExtendedPrice         float64 `json:"extendedPrice"`
		ExtendedChange        float64 `json:"extendedChange"`
		ExtendedChangePercent float64 `json:"extendedChangePercent"`
		ExtendedPriceTime     int64   `json:"extendedPriceTime"`
		PreviousClose         float64 `json:"previousClose"`
		Change                float64 `json:"change"`
		ChangePercent         float64 `json:"changePercent"`
		IexMarketPercent      float64 `json:"iexMarketPercent"`
		IexVolume             int     `json:"iexVolume"`
		AvgTotalVolume        int     `json:"avgTotalVolume"`
		IexBidPrice           int     `json:"iexBidPrice"`
		IexBidSize            int     `json:"iexBidSize"`
		IexAskPrice           int     `json:"iexAskPrice"`
		IexAskSize            int     `json:"iexAskSize"`
		MarketCap             int64   `json:"marketCap"`
		PeRatio               float64 `json:"peRatio"`
		Week52High            float64 `json:"week52High"`
		Week52Low             float64 `json:"week52Low"`
		YtdChange             float64 `json:"ytdChange"`
	} `json:"quote"`
	News []struct {
		Datetime string `json:"datetime"`
		Headline string `json:"headline"`
		Source   string `json:"source"`
		URL      string `json:"url"`
		Summary  string `json:"summary"`
		Related  string `json:"related"`
		Image    string `json:"image"`
	} `json:"news"`
	Chart []struct {
		Date             string  `json:"date"`
		Open             float64 `json:"open"`
		High             float64 `json:"high"`
		Low              float64 `json:"low"`
		Close            float64 `json:"close"`
		Volume           int     `json:"volume"`
		UnadjustedVolume int     `json:"unadjustedVolume"`
		Change           float64 `json:"change"`
		ChangePercent    float64 `json:"changePercent"`
		Vwap             float64 `json:"vwap"`
		Label            string  `json:"label"`
		ChangeOverTime   float64 `json:"changeOverTime"`
	} `json:"chart"`
}

func (stocks StockList) String() string {
	var list []string
	for _, s := range stocks {
		list = append(list, s.Quote.Symbol)
	}
	final := strings.Join(list, ",")

	return final
}

func (stock Stock) GetDatesChart() []time.Time {
	var finalSlice []time.Time
	layout := "2006-01-02"
	for _, v := range stock.Chart {
		date, err := time.Parse(layout, v.Date)
		if err != nil {
			fmt.Println("Error parsing date string", err)
		}
		finalSlice = append(finalSlice, date)
	}
	return finalSlice
}

func (stock Stock) GetFloatChart() []float64 {
	var finalSlice []float64

	for _, v := range stock.Chart {

		finalSlice = append(finalSlice, v.Close)
	}
	return finalSlice
}

// func DateStringToInt(date string) int {

// }
