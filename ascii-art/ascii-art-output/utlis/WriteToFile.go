package ascii

import (
	"os"
)

func WriteToFile(fileName string, data string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write to the file
	_, err = file.WriteString(data)

	return err
}
