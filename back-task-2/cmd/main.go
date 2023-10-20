package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const URL = "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1"

type CoinData struct {
	Symbol       string  `json:"symbol"`
	Name         string  `json:"name"`
	CurrentPrice float64 `json:"current_price"`
}

func fetchCoinData() ([]CoinData, error) {

	var client = &http.Client{Timeout: 2 * time.Second}

	resp, err := client.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed, status: %s", resp.Status)
	}

	var data []CoinData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func main() {
	coinSymbol := "btc" // change coinSymbol

	timeInterval := 10 * time.Minute

	cryptoData, err := fetchCoinData()
	if err != nil {
		fmt.Printf("failed to fetch data: %v\n", err)
		return
	}

	var found bool

	for _, coin := range cryptoData {
		if coin.Symbol == coinSymbol {
			fmt.Printf("Course of currency %v (%v) is: %v", coin.Name, coin.Symbol, coin.CurrentPrice)
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("currency with symbol %s was not fiund", coinSymbol)
	}

	time.Sleep(timeInterval)
}
