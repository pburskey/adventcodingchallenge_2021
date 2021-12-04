package utility

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func assembleFilePathToDay(day string) string {
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

func ParseInputFileIntoStringRows(path string) ([]string, error) {
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

func parseInputFileIntoNumberRows(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data := make([]int, 0)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for i := 0; fileScanner.Scan(); i++ {
		aString := fileScanner.Text()
		aNumber, err := strconv.Atoi(aString)
		if err == nil {
			log.Fatalln(err)
		}
		data = append(data, aNumber)
	}
	file.Close()
	return data, nil
}

func ParseDayForInputIntoStringRows(day string, path string) ([]string, error) {
	return ParseInputFileIntoStringRows(filepath.Join(assembleFilePathToDay(day), path))
}

func ParseDayForInputIntoNumberRows(day string, path string) ([]int, error) {
	return parseInputFileIntoNumberRows(filepath.Join(assembleFilePathToDay(day), path))
}
