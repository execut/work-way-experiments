//go:generate mockgen -source service.go -destination service_mock.go -package order order Service
package order

type Service interface {
    GetTotalOrdersSum() uint64
}

func New() Service {
    return &Order{}
}

type Order struct {
}

func (s *Order) GetTotalOrdersSum() uint64 {
    return 5000
}
