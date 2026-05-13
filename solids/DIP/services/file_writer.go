package services

import "os"

// FileWriter is a concrete implementation of the Writer interface that writes data to a file.
type FileWriter struct {
	FilePath string //path to the file where data will be written
}

// Write writes the given data to the specified file path.
func (fw *FileWriter) Write(data []byte) error {
	// Logic to write data to a file at fw.FilePath
	file, err := os.Create(fw.FilePath) //create or overwrite the file at the specified path
	if err != nil {
		return err //return error if file creation fails
	}
	defer file.Close()
	_, err = file.Write(data) //write data to the file
	if err != nil {
		return err //return error if writing to the file fails
	}
	return nil
}
