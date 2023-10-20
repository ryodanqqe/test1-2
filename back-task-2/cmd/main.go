package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const coingeckoURL = "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1"

type CoinData struct {
	ID           string  `json:"id"`
	Symbol       string  `json:"symbol"`
	Name         string  `json:"name"`
	CurrentPrice float64 `json:"current_price"`
}

var (
	client         = &http.Client{Timeout: 10 * time.Second}
	cache          = make(map[string]CoinData)
	lastUpdateTime time.Time
)

func fetchCoinData() error {
	resp, err := client.Get(coingeckoURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status: %s", resp.Status)
	}

	var data []CoinData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return err
	}

	// Update cache and lastUpdateTime
	for _, coin := range data {
		cache[coin.Symbol] = coin
	}
	lastUpdateTime = time.Now()

	return nil
}

func getCoinPrice(symbol string) (float64, error) {
	if time.Since(lastUpdateTime) > 10*time.Minute {
		if err := fetchCoinData(); err != nil {
			return 0, err
		}
	}

	if coin, found := cache[symbol]; found {
		return coin.CurrentPrice, nil
	}

	return 0, fmt.Errorf("cryptocurrency with symbol %s not found", symbol)
}

func main() {
	// Fetch initial data
	if err := fetchCoinData(); err != nil {
		fmt.Printf("Failed to fetch initial data: %v\n", err)
		return
	}

	// Example: Get the price of Bitcoin (BTC)
	price, err := getCoinPrice("btc")
	if err != nil {
		fmt.Printf("Failed to get price: %v\n", err)
		return
	}

	fmt.Printf("Price of BTC: $%.2f USD\n", price)
}
