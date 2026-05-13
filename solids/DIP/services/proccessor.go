package services

import "system-design/solids/DIP/models"

// Processor is a high-level module that depends on the Writer interface.
type Processor struct {
	writer models.Writer //writer is an abstraction that Processor depends on
}

// NewProcessor creates a new instance of Processor with the given Writer.
func NewProcessor(writer models.Writer) *Processor {
	return &Processor{writer: writer}
}

// ProcessData processes the given data and writes it using the injected Writer.
func (p *Processor) ProcessData(data []byte) error {
	// Logic to process data before writing
	println("Processing data...")
	return p.writer.Write(data) //write processed data using the injected Writer
}
