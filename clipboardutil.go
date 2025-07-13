package main

import (
	"github.com/atotto/clipboard"
)

func ReadFromClipboard() (string, error) {
	return clipboard.ReadAll()
}

func writeToClipboard(content string) error {
	return clipboard.WriteAll(content)
}
