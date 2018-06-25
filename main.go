package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/firstimedeveloper/finance/auth"

	"github.com/wcharczuk/go-chart"

	"github.com/firstimedeveloper/finance/stock"
)

func main() {
	url := "https://api.iextrading.com/1.0/stock/aapl/batch?types=quote,news,chart&range=1y&last=1"
	///stock/market/batch?symbols=aapl,fb,tsla&types=quote,news,chart&range=1m&last=5
	///stock/aapl/batch?types=quote,news,chart&range=1m&last=1

	//textJson := `{"quote":{"symbol":"AAPL","companyName":"Apple Inc.","primaryExchange":"Nasdaq Global Select","sector":"Technology","calculationPrice":"close","open":185.94,"openTime":1529674200375,"close":184.92,"closeTime":1529697600451,"high":186.15,"low":184.7,"latestPrice":184.92,"latestSource":"Close","latestTime":"June 22, 2018","latestUpdate":1529697600451,"latestVolume":27254380,"iexRealtimePrice":185.39,"iexRealtimeSize":5,"iexLastUpdated":1529697668550,"delayedPrice":184.92,"delayedPriceTime":1529697600451,"extendedPrice":185.01,"extendedChange":0.09,"extendedChangePercent":0.00049,"extendedPriceTime":1529701087475,"previousClose":185.46,"change":-0.54,"changePercent":-0.00291,"iexMarketPercent":0.01598,"iexVolume":435525,"avgTotalVolume":23148364,"iexBidPrice":0,"iexBidSize":0,"iexAskPrice":0,"iexAskSize":0,"marketCap":908907318960,"peRatio":19.01,"week52High":194.2,"week52Low":142.28,"ytdChange":0.07911678150864673},"news":[{"datetime":"2018-06-23T12:29:00-04:00","headline":"How to choose between an Amazon Echo, a Google Home and an Apple HomePod","source":"CNBC","url":"https://api.iextrading.com/1.0/stock/aapl/article/8072348731447178","summary":"No summary available.","related":"AAPL,AMZN,GOOGL","image":"https://api.iextrading.com/1.0/stock/aapl/news-image/8072348731447178"}],"chart":[{"date":"2018-05-23","open":186.35,"high":188.5,"low":185.76,"close":188.36,"volume":20058415,"unadjustedVolume":20058415,"change":1.2,"changePercent":0.641,"vwap":187.0702,"label":"May 23","changeOverTime":0},{"date":"2018-05-24","open":188.77,"high":188.84,"low":186.21,"close":188.15,"volume":23233975,"unadjustedVolume":23233975,"change":-0.21,"changePercent":-0.111,"vwap":187.7764,"label":"May 24","changeOverTime":-0.0011148863877681458},{"date":"2018-05-25","open":188.23,"high":189.65,"low":187.65,"close":188.58,"volume":17460963,"unadjustedVolume":17460963,"change":0.43,"changePercent":0.229,"vwap":188.8718,"label":"May 25","changeOverTime":0.0011679762157570548},{"date":"2018-05-29","open":187.6,"high":188.75,"low":186.87,"close":187.9,"volume":22514075,"unadjustedVolume":22514075,"change":-0.68,"changePercent":-0.361,"vwap":187.6739,"label":"May 29","changeOverTime":-0.0024421320874920788},{"date":"2018-05-30","open":187.72,"high":188,"low":186.78,"close":187.5,"volume":18690547,"unadjustedVolume":18690547,"change":-0.4,"changePercent":-0.213,"vwap":187.2798,"label":"May 30","changeOverTime":-0.004565725207050402},{"date":"2018-05-31","open":187.22,"high":188.23,"low":186.14,"close":186.87,"volume":27482793,"unadjustedVolume":27482793,"change":-0.63,"changePercent":-0.336,"vwap":187.123,"label":"May 31","changeOverTime":-0.007910384370354687},{"date":"2018-06-01","open":187.9912,"high":190.26,"low":187.75,"close":190.24,"volume":23442510,"unadjustedVolume":23442510,"change":3.37,"changePercent":1.803,"vwap":189.4938,"label":"Jun 1","changeOverTime":0.00998088766192395},{"date":"2018-06-04","open":191.635,"high":193.42,"low":191.35,"close":191.83,"volume":26266174,"unadjustedVolume":26266174,"change":1.59,"changePercent":0.836,"vwap":192.2436,"label":"Jun 4","changeOverTime":0.018422170312168182},{"date":"2018-06-05","open":193.065,"high":193.94,"low":192.36,"close":193.31,"volume":21565963,"unadjustedVolume":21565963,"change":1.48,"changePercent":0.772,"vwap":193.1978,"label":"Jun 5","changeOverTime":0.02627946485453381},{"date":"2018-06-06","open":193.63,"high":194.08,"low":191.92,"close":193.98,"volume":20933619,"unadjustedVolume":20933619,"change":0.67,"changePercent":0.347,"vwap":193.2484,"label":"Jun 6","changeOverTime":0.029836483329793884},{"date":"2018-06-07","open":194.14,"high":194.2,"low":192.335,"close":193.46,"volume":21347180,"unadjustedVolume":21347180,"change":-0.52,"changePercent":-0.268,"vwap":193.2897,"label":"Jun 7","changeOverTime":0.0270758122743682},{"date":"2018-06-08","open":191.17,"high":192,"low":189.77,"close":191.7,"volume":26656799,"unadjustedVolume":26656799,"change":-1.76,"changePercent":-0.91,"vwap":190.9465,"label":"Jun 8","changeOverTime":0.017732002548311608},{"date":"2018-06-11","open":191.35,"high":191.97,"low":190.21,"close":191.23,"volume":18308460,"unadjustedVolume":18308460,"change":-0.47,"changePercent":-0.245,"vwap":191.1869,"label":"Jun 11","changeOverTime":0.015236780632830621},{"date":"2018-06-12","open":191.385,"high":192.611,"low":191.15,"close":192.28,"volume":16911141,"unadjustedVolume":16911141,"change":1.05,"changePercent":0.549,"vwap":191.7894,"label":"Jun 12","changeOverTime":0.0208112125716712},{"date":"2018-06-13","open":192.42,"high":192.88,"low":190.44,"close":190.7,"volume":21638393,"unadjustedVolume":21638393,"change":-1.58,"changePercent":-0.822,"vwap":191.5742,"label":"Jun 13","changeOverTime":0.012423019749415879},{"date":"2018-06-14","open":191.55,"high":191.57,"low":190.22,"close":190.8,"volume":21610074,"unadjustedVolume":21610074,"change":0.1,"changePercent":0.052,"vwap":190.8927,"label":"Jun 14","changeOverTime":0.012953918029305572},{"date":"2018-06-15","open":190.03,"high":190.16,"low":188.26,"close":188.84,"volume":61719160,"unadjustedVolume":61719160,"change":-1.96,"changePercent":-1.027,"vwap":189.0795,"label":"Jun 15","changeOverTime":0.0025483117434698965},{"date":"2018-06-18","open":187.88,"high":189.22,"low":187.2,"close":188.74,"volume":18484865,"unadjustedVolume":18484865,"change":-0.1,"changePercent":-0.053,"vwap":188.6013,"label":"Jun 18","changeOverTime":0.002017413463580354},{"date":"2018-06-19","open":185.14,"high":186.33,"low":183.45,"close":185.69,"volume":33578455,"unadjustedVolume":33578455,"change":-3.05,"changePercent":-1.616,"vwap":185.1633,"label":"Jun 19","changeOverTime":-0.014174984073051686},{"date":"2018-06-20","open":186.35,"high":187.2,"low":185.73,"close":186.5,"volume":20628701,"unadjustedVolume":20628701,"change":0.81,"changePercent":0.436,"vwap":186.5042,"label":"Jun 20","changeOverTime":-0.009874708005946132},{"date":"2018-06-21","open":187.25,"high":188.35,"low":184.94,"close":185.46,"volume":25711898,"unadjustedVolume":25711898,"change":-1.04,"changePercent":-0.558,"vwap":186.1556,"label":"Jun 21","changeOverTime":-0.015396050116797651},{"date":"2018-06-22","open":186.12,"high":186.15,"low":184.7,"close":184.92,"volume":27200447,"unadjustedVolume":27200447,"change":-0.54,"changePercent":-0.291,"vwap":185.291,"label":"Jun 22","changeOverTime":-0.018262900828201454}]}`

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Problem with fetching data")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	newStock := UnmarshalJson(body)

	dateSlice := newStock.GetDatesChart()

	fmt.Println(newStock.Quote.CompanyName)
	fmt.Println(dateSlice)

	graph := chart.Chart{
		Series: []chart.Series{
			chart.TimeSeries{
				XValues: dateSlice,
				YValues: newStock.GetFloatChart(),
			},
		},
	}

	f, _ := os.Create("image.png")

	//buffer := bytes.NewBuffer([]byte{})
	err = graph.Render(chart.PNG, f)
	if err != nil {
		fmt.Println("Error creating graph", err)
	}

	auth.Authenticate()

}

func UnmarshalJson(textBytes []byte) stock.Stock {
	//textBytes := []byte(textJson)
	newStock := stock.Stock{}
	err := json.Unmarshal(textBytes, &newStock)
	if err != nil {
		fmt.Println("Unable to unmarshall json", err)
	}
	return newStock
}
