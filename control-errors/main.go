package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type NoTxtError struct {
	Message string
}

func (n *NoTxtError) Error() string {
	return n.Message
}

func loadFile(filename string) (string, error) {
	if !strings.HasSuffix(filename, ".txt") {
		return "", &NoTxtError{
			"Opening non txt files not allowed",
		}
	}

	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	return filename, nil
}

func main() {
	s, err := loadFile("go.mod")

	var noTxtError *NoTxtError
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File does not exist")
	} else if errors.As(err, &noTxtError) {
		fmt.Println("Txt error:", noTxtError.Message)
	}

	fmt.Println("The string:", s)
}
