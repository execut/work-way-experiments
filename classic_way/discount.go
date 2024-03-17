package classic_way

func New() *Service {
    return &Service{}
}

type Service struct {
}

func (s *Service) CalculateDiscount(price int64) (int64, error) {
    return 0, nil
}
