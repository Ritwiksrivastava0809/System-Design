package models

//LSP models define the structures and interfaces for the Liskov Substitution Principle example.

type PaymentMethod interface {
	ProcessPayment(amount float64) error //process payment with the given amount
	Name() string                        //return the name of the payment method
}
