package services

import "fmt"

// DocumentScanner only supports scanning.
type DocumentScanner struct{}

// Scan scans the document.
func (s *DocumentScanner) Scan(
	document string,
) error {

	fmt.Printf(
		"[SCANNER] Scanning: %s\n",
		document,
	)

	return nil
}
