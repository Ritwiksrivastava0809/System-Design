package models

// ================================================================
// Interface Segregation Principle (ISP)
// ================================================================
//
// ISP:
// Clients should not depend on methods
// they do not use.
//
// Small, focused interfaces are preferred in Go.
// ================================================================

// Printer represents printing capability.
type Printer interface {
	Print(document string) error
}

// Scanner represents scanning capability.
type Scanner interface {
	Scan(document string) error
}

// Fax represents fax capability.
type Fax interface {
	SendFax(document string) error
}

// ================================================================
// BAD DESIGN — Violates ISP
// ================================================================
//
// A client implementing this interface is forced
// to implement ALL methods even if it only needs one.
//

type MultiFunctionDevice interface {
	Print(document string) error
	Scan(document string) error
	SendFax(document string) error
}
