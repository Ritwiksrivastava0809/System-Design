package main

import (
	"system-design/solids/DIP/services"
)

func main() {
	// The main function is intentionally left empty as the focus is on the DIP example.

	fw := services.FileWriter{
		FilePath: "output.txt",
	}
	processor := services.NewProcessor(&fw) // Injecting the FileWriter dependency into the Processor.

	data := []byte("Hello, Dependency Inversion Principle!")
	err := processor.ProcessData(data) // Processing data and writing it using the injected FileWriter.
	if err != nil {
		println("Error processing data:", err.Error())
	} else {
		println("Data processed and written successfully.")
	}

	nw := services.NetworkWriter{
		Address: "www.example.com",
	}
	processor = services.NewProcessor(&nw) // Injecting the NetworkWriter dependency into the Processor.

	err = processor.ProcessData(data) // Processing data and writing it using the injected NetworkWriter.
	if err != nil {
		println("Error processing data:", err.Error())
	} else {
		println("Data processed and written to network successfully.")
	}

}
