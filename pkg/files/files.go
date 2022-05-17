package files

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Errors

var (
	FileNotFoundError    func(string) error = func(file string) error { return fmt.Errorf("file not found : %s", file) }
	CannotCloseFileError func(string) error = func(file string) error { return fmt.Errorf("the application cannot close with this file : %s", file) }
)

func Read(data io.Reader) string {

	scanner := bufio.NewScanner(data)
	result := []string{}
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return strings.Join(result, "\n")
}
