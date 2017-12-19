package services

import (
	"fmt"
	"time"
	"github.com/Junbong/coinsightly-api/models"
	"net/http"
	"log"
	"io/ioutil"
	"database/sql"
	_ "github.com/lib/pq"
)

const (
	apiFetchTimeGap = time.Second * 10
	apiFetchTimeout = time.Second * 10
)

func RunFetcher() {
	connectDb()

	markets := []models.Market{
		{
			Market: "bitfinex",
			CoinSct: "btc",
			Currency: "usd",
		},
		{
			Market: "bithumb",
			CoinSct: "btc",
			Currency: "krw",
		},
	}

	go fetchLooper(markets)
}

func connectDb() {
	connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"
	_, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

func fetchLooper(markets []models.Market) {
	for {
		for _, market := range markets {
			go fetch(market)
		}
		time.Sleep(apiFetchTimeGap)
	}
}

func fetch(market models.Market) {
	url := fmt.Sprintf("https://api.cryptowat.ch/markets/%s/%s%s/ohlc",
		market.Market, market.CoinSct, market.Currency)

	httpClient := http.Client{
		Timeout: apiFetchTimeout,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	_, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	log.Println("Fetch", market.Market, market.CoinSct, url)
}
