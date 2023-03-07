package main

import (
	"fmt"

	"github.com/google/uuid"
)

type Order struct {
	ID         string
	Price      float64
	Quantidade int
}

func (o *Order) SetPrice(price float64) {
	o.Price = price
	fmt.Println("Price dentro do SetPrice: ", o.Price)
}

func (o Order) getTotal() float64 {
	return o.Price * float64(o.Quantidade)
}

func NewOrder() Order {
	return Order{}
}

func main() {
	fmt.Println(uuid.New().String())
}
