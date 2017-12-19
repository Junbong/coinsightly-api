package models

type Market struct {
	Market string
	CoinSct string
	Currency string
}

type Coin struct {
	TimeKey string
	CloseTime uint16
	OpenPrice uint8
	HighPrice uint8
	LowPrice uint8
	ClosePrice float32
	Volume float64
}
