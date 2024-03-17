package tdd_way

import (
    "errors"
    "go.uber.org/mock/gomock"
    "tdd_experiment/order"
    "testing"
)

func TestService_CalculateDiscount(t *testing.T) {
    type fields struct {
        sum uint64
    }
    type args struct {
        price int64
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        want    int64
        wantErr error
    }{
        {
            name: "посчитать скидку от переданной суммы в размере 10%",
            fields: fields{
                sum: 10000,
            },
            args: args{
                price: 100,
            },
            want: 10,
        },
        {
            name: "с округлением в меньшую сторону (109)",
            fields: fields{
                sum: 10000,
            },
            args: args{
                price: 109,
            },
            want: 10,
        },
        {
            name: "сумма не может быть отрицательной или равна нулю",
            fields: fields{
                sum: 10000,
            },
            args: args{
                price: -1,
            },
            wantErr: errLessThanZero,
        },
        {
            name: "сумма не может превышать 1 миллиард (1000000000)",
            fields: fields{
                sum: 10000,
            },
            args: args{
                price: 1000000001,
            },
            wantErr: errMoreThanMilliard,
        },
        {
            name: "0% если сумма заказов <2000",
            fields: fields{
                sum: 1999,
            },
            args: args{
                price: 100,
            },
            want: 0,
        },
        {
            name: "3% если сумма заказов >=2000 <10000",
            fields: fields{
                sum: 2000,
            },
            args: args{
                price: 100,
            },
            want: 3,
        },
        {
            name: "5% если сумма заказов >=5000 <10000",
            fields: fields{
                sum: 5000,
            },
            args: args{
                price: 100,
            },
            want: 5,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            ctrl := gomock.NewController(t)
            mockOrders := order.NewMockService(ctrl)
            mockOrders.EXPECT().GetTotalOrdersSum().AnyTimes().Return(tt.fields.sum)

            s := &Service{
                service: mockOrders,
            }

            got, err := s.CalculateDiscount(tt.args.price)

            if tt.wantErr != nil && !errors.Is(err, tt.wantErr) {
                t.Errorf("CalculateDiscount() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("CalculateDiscount() got = %v, want %v", got, tt.want)
            }
        })
    }
}
