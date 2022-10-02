package gas

type RequestData struct {
	Ethereum Ethereum `json:"ethereum"`
}

type Ethereum struct {
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	Time           string  `json:"time"`
	GasPrice       float64 `json:"gasPrice"`
	GasValue       float64 `json:"gasValue"`
	Average        float64 `json:"average"`
	MaxGasPrice    float64 `json:"maxGasPrice"`
	MedianGasPrice float64 `json:"medianGasPrice"`
}

type ResponseData struct {
	SpentPerMonth         map[string]float64      `json:"spentPerMonth"`
	TotalCost             float64                 `json:"totalСosts"`
	AveragePricePerDay    map[string]float64      `json:"averagePricePerDay"`
	FrequencyDistribution map[string]Distribution `json:"frequencyDistribution"`
}

type Distribution struct {
	Max float64 `json:"max"`
	Min float64 `json:"min"`
}
