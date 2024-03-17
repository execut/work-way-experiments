package main

import (
    "fmt"
    "tdd_experiment/order"

    //oldWay "tdd_experiment/classic_way"
    tddWay "tdd_experiment/tdd_way"
)

func NewService() IService {
    return tddWay.New(order.New())
}

type IService interface {
    CalculateDiscount(price int64) (int64, error)
}

func main() {
    var price int64
    fmt.Printf("Enter price: ")
    _, err := fmt.Scanf("%d", &price)
    if err != nil {
        panic(err)
    }

    service := NewService()
    discount, err := service.CalculateDiscount(price)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Discount price: %d\n", discount)
}
