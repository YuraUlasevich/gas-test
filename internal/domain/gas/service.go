package gas

import "gas-test/internal/adapters/gas"

type service struct {
}

func NewService() gas.Service {
	return &service{}
}

func (s *service) FrequencyDistribution() {

}

func (s *service) AveragePricePerDay() {

}

func (s *service) SpentInMonth() {

}

func (s *service) TotalСosts() {

}
