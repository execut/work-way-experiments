package tdd_way

import (
    "errors"
    "tdd_experiment/order"
)

var (
    errLessThanZero     = errors.New("сумма не может быть отрицательной или равна нулю")
    errMoreThanMilliard = errors.New("сумма не может превышать 1 миллиард (1000000000)")
)

func New(service order.Service) *Service {
    return &Service{
        service: service,
    }
}

type Service struct {
    service order.Service
}

func (s *Service) CalculateDiscount(price int64) (int64, error) {
    if price <= 0 {
        return 0, errLessThanZero
    }

    if price > 1000000000 {
        return 0, errMoreThanMilliard
    }

    sum := s.service.GetTotalOrdersSum()
    if sum < 2000 {
        return 0, nil
    }

    if sum >= 2000 && sum < 5000 {
        return 3, nil
    }

    if sum >= 5000 && sum < 10000 {
        return 5, nil
    }

    return price / 10, nil
}
