package gas

import "sync"

type Service interface {
	FrequencyDistribution(wg *sync.WaitGroup, gasData *RequestData, result *ResponseData)
	AveragePricePerDay(wg *sync.WaitGroup, gasData *RequestData, result *ResponseData)
	SpentPerMonth(wg *sync.WaitGroup, gasData *RequestData, result *ResponseData)
	TotalСost(wg *sync.WaitGroup, gasData *RequestData, result *ResponseData)
}
