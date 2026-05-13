package services

import "fmt"

// BasicPrinter only supports printing.
type BasicPrinter struct{}

// Print prints the document.
func (p *BasicPrinter) Print(
	document string,
) error {

	fmt.Printf(
		"[PRINTER] Printing: %s\n",
		document,
	)

	return nil
}
