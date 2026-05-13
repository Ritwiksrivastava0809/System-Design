package services

import "fmt"

// MultiFunctionPrinter supports all operations.
type MultiFunctionPrinter struct{}

// Print prints the document.
func (m *MultiFunctionPrinter) Print(
	document string,
) error {

	fmt.Printf(
		"[MFP] Printing: %s\n",
		document,
	)

	return nil
}

// Scan scans the document.
func (m *MultiFunctionPrinter) Scan(
	document string,
) error {

	fmt.Printf(
		"[MFP] Scanning: %s\n",
		document,
	)

	return nil
}

// SendFax sends the fax.
func (m *MultiFunctionPrinter) SendFax(
	document string,
) error {

	fmt.Printf(
		"[MFP] Faxing: %s\n",
		document,
	)

	return nil
}
