package services

import "fmt"

// FaxMachine only supports faxing.
type FaxMachine struct{}

// SendFax sends the fax document.
func (f *FaxMachine) SendFax(
	document string,
) error {

	fmt.Printf(
		"[FAX] Sending fax: %s\n",
		document,
	)

	return nil
}
