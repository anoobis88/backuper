package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Define the source and destination paths
	sourcePath := "z:\\bp_absol\\"
	destinationPath := "/path/to/destination/file.txt"

	// Open the source file
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		fmt.Println("Failed to open source file:", err)
		return
	}
	defer func(sourceFile *os.File) {
		err := sourceFile.Close()
		if err != nil {

		}
	}(sourceFile)

	// Create the destination file
	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		fmt.Println("Failed to create destination file:", err)
		return
	}
	defer func(destinationFile *os.File) {
		err := destinationFile.Close()
		if err != nil {

		}
	}(destinationFile)

	// Copy the contents from source to destination
	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		fmt.Println("Failed to copy file:", err)
		return
	}

	fmt.Println("File copied successfully!")
}
