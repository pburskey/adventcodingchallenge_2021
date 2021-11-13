package utility

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func AssembleFilePathToDay(day string) string {
	path, _ := os.Getwd()
	dayMunge := fmt.Sprintf("day%s", day)

	/*
		necessary to conveniently correct working directory issues present in goland vs terminal.
	*/
	if !strings.Contains(path, dayMunge) {
		path = filepath.Join(path, dayMunge)
	}
	return path
}

func ParseInputFileIntoRows(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data := make([]string, 0)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for i := 0; fileScanner.Scan(); i++ {
		aString := fileScanner.Text()
		data = append(data, aString)

	}

	file.Close()

	return data, nil
}
