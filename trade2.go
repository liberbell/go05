package main

type Trade struct {
	Symbol string
	Volume int
	Price  float64
	buy    bool
}

func NewTrade(symbol string, volume int, price float64, buy bool) (+Trade, error) {
  if symbol == "" {
    return nil, fmt.Errorf("symbol can`t be empty.")
  }
}
