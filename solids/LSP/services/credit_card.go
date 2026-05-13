package services

type CreditCardPayment struct{}

// ProcessPayment processes the payment using credit card.
func (ccp *CreditCardPayment) ProcessPayment(amount float64) error {
	// Logic to process credit card payment
	println("Processing credit card payment of amount:", amount)
	return nil
}

func (ccp *CreditCardPayment) Name() string {
	return "Credit Card"
}
