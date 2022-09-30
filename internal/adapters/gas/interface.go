package gas

type Service interface {
	FrequencyDistribution()
	AveragePricePerDay()
	SpentInMonth()
	TotalСosts()
}
