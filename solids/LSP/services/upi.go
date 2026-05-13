package services

type UPIPayment struct{}

// ProcessPayment processes the payment using UPI.
func (upip *UPIPayment) ProcessPayment(amount float64) error {
	// Logic to process UPI payment
	println("Processing UPI payment of amount:", amount)
	return nil
}

func (upip *UPIPayment) Name() string {
	return "UPI"
}
