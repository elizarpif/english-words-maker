package maker

import (
	"os"
	"strings"
)

// writeToFile writes []string to file
func writeToFile(values []string, fileName string) error {
	newWordsFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer newWordsFile.Close()

	_, err = newWordsFile.WriteString(strings.Join(values, "\n"))
	return err
}
