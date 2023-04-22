package main

import (
	"fmt"
	"math/rand"
)

type FeePayer struct {
	NumOfEmployees  int
	PaymentStrategy PaymentStrategy
	Budget          float64
}

func (f *FeePayer) PaySalary() {
	for i := 0; i < f.NumOfEmployees; i++ {
		fmt.Printf("Paying salary for %d employee\n", i+1)
		fmt.Println("Payment:", f.PaymentStrategy.CalculatePayment())
		fmt.Println()
	}
}

type PaymentStrategy interface {
	CalculatePayment() float64
}

type DefaultPayment struct {
}

func (d *DefaultPayment) CalculatePayment() float64 {
	return rand.Float64()*9000 + 1000
}

type BonusPayment struct {
	BonusSize float64
}

func (b *BonusPayment) CalculatePayment() float64 {
	return rand.Float64()*9000 + 1000 + b.BonusSize
}

func main() {
	f := FeePayer{
		NumOfEmployees:  10,
		PaymentStrategy: &DefaultPayment{},
		Budget:          250_000,
	}

	f.PaySalary()

	f.PaymentStrategy = &BonusPayment{BonusSize: 5_000}

	fmt.Println("With bonus")
	f.PaySalary()
}
