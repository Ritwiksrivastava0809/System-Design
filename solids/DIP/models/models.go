package models

// writer interface defines the contract for writing data to a destination.
type Writer interface {
	Write(data []byte) error //write data to the destination
}
