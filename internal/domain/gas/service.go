package gas

import (
	"fmt"
	"gas-test/internal/adapters/gas"
	"strings"
	"sync"
)

type service struct {
}

func NewService() gas.Service {
	return &service{}
}

func (s *service) FrequencyDistribution(wg *sync.WaitGroup, gasData *gas.RequestData, result *gas.ResponseData) {
	defer wg.Done()

	freqDistr := make(map[string]gas.Distribution, 24)
	for _, v := range gasData.Ethereum.Transactions {
		time := strings.Split(v.Time, " ")[1]
		hour := strings.Split(time, ":")[0]
		if _, ok := freqDistr[hour]; !ok {
			freqDistr[hour] = gas.Distribution{
				Max: 0,
				Min: v.MaxGasPrice,
			}
		}

		if freqDistr[hour].Max < v.GasPrice {
			freqDistr[hour] = gas.Distribution{
				Max: v.GasPrice,
				Min: freqDistr[hour].Min,
			}
		}
		if freqDistr[hour].Min > v.GasPrice {
			freqDistr[hour] = gas.Distribution{
				Max: freqDistr[hour].Max,
				Min: v.GasPrice,
			}
		}
	}

	result.FrequencyDistribution = freqDistr
}

func (s *service) AveragePricePerDay(wg *sync.WaitGroup, gasData *gas.RequestData, result *gas.ResponseData) {
	defer wg.Done()

	dayStat := make(map[string]DayStat)
	for _, v := range gasData.Ethereum.Transactions {
		date := strings.Split(v.Time, " ")[0]
		dayStat[date] = DayStat{
			Sum:   dayStat[date].Sum + v.GasPrice,
			Count: dayStat[date].Count + 1,
		}
	}

	an := make(map[string]float64)
	for k, v := range dayStat {
		an[k] = v.Sum / v.Count
	}

	result.AveragePricePerDay = an
}

func (s *service) SpentPerMonth(wg *sync.WaitGroup, gasData *gas.RequestData, result *gas.ResponseData) {
	defer wg.Done()

	spent := make(map[string]float64)
	for _, v := range gasData.Ethereum.Transactions {
		s := strings.Split(v.Time, "-")[0:2]
		month := fmt.Sprintf("%s-%s", s[0], s[1])
		spent[month] += v.GasValue
	}

	result.SpentPerMonth = spent
}

func (s *service) TotalCost(wg *sync.WaitGroup, gasData *gas.RequestData, result *gas.ResponseData) {
	defer wg.Done()

	var total float64
	for _, v := range gasData.Ethereum.Transactions {
		total += v.GasPrice * v.GasValue
	}

	result.TotalCost = total
}
