package utils

import (
	"bufio"
	"fmt"
	"os"
)

func CreateFile(fileName string) {
	if FileExists(fileName) {
		return
	}

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer file.Close()
}

func FileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil || !os.IsNotExist(err)
}

func ReadFileToArray(fileName string) []string {
	file, err := ReadRawFile(fileName)

	if err != nil {
		return nil
	}

	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	return lines
}

func ReadRawFile(fileName string) (*os.File, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	return file, nil
}
