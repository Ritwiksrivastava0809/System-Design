package main

import (
	"fmt"
	"system-design/solids/LSP/models"
	"system-design/solids/LSP/services"
)

// LSP main function demonstrates the Liskov Substitution
// Principle by processing payments using different payment methods that implement the same interface.
// We can substitute one payment method with another without affecting the correctness of the program.
// In this example, both UPI and Credit Card payment methods implement the PaymentMethod interface,
// allowing us to process payments using either method seamlessly.
func main() {
	// Create instances of different payment methods that implement the PaymentMethod interface.
	upiPayment := &services.UPIPayment{}
	creditCardPayment := &services.CreditCardPayment{}
	//LSP models define the structures and interfaces for the Liskov Substitution Principle example.
	paymentMethods := []models.PaymentMethod{upiPayment, creditCardPayment}
	// Process payments using the different payment methods without worrying about their specific implementations.
	// We can substitute one payment method with another without affecting the correctness of the program.
	for _, paymentMethod := range paymentMethods {
		processor := services.NewTransactionProcessor(paymentMethod)

		err := processor.ProcessTransaction(100)
		if err != nil {
			println(err.Error())
		}
		fmt.Printf("successfull payment by provider :: %v\n", paymentMethod.Name())
	}
}
