package main

import (
	"system-design/solids/ISP/models"
	"system-design/solids/ISP/services"
)

// ================================================================
// ISP Main
// ================================================================
//
// Demonstrates:
// - small focused interfaces
// - clients depending only on required behavior
// - better modularity and flexibility
//

func main() {

	// ============================================================
	// Printer Example
	// ============================================================

	var printer models.Printer = &services.BasicPrinter{}

	printer.Print("invoice.pdf")

	// ============================================================
	// Scanner Example
	// ============================================================

	var scanner models.Scanner = &services.DocumentScanner{}

	scanner.Scan("contract.pdf")

	// ============================================================
	// Fax Example
	// ============================================================

	var fax models.Fax = &services.FaxMachine{}

	fax.SendFax("legal_notice.pdf")

	// ============================================================
	// Multifunction Printer Example
	// ============================================================

	mfp := &services.MultiFunctionPrinter{}

	mfp.Print("report.pdf")
	mfp.Scan("passport.pdf")
	mfp.SendFax("agreement.pdf")
}
