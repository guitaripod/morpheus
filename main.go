package main

import (
	"fmt"
	"strings"
)

func main() {
	clipboardContent, err := ReadFromClipboard()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	imgFilePath := strings.TrimSpace(clipboardContent)
	if imgFilePath == "" {
		fmt.Println("Error: clipboard content is empty")
		return
	}

	outputPath, err := ProcessImage(imgFilePath, "/Users/marcusziade/Downloads")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	err = writeToClipboard(outputPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Image saved to %s\n", outputPath)
}
