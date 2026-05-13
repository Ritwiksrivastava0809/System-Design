package services

import "system-design/solids/LSP/models"

type TransactionProcessor struct {
	paymentMethod models.PaymentMethod
}

// ProcessTransaction processes the transaction using the provided payment method.
func (tp *TransactionProcessor) ProcessTransaction(amount float64) error {
	return tp.paymentMethod.ProcessPayment(amount)
}

// NewTransactionProcessor creates a new TransactionProcessor with the given payment method.
func NewTransactionProcessor(paymentMethod models.PaymentMethod) *TransactionProcessor {
	return &TransactionProcessor{
		paymentMethod: paymentMethod,
	}
}
