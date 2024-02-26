package main

import (
	"encoding/csv"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type readFunc func(filepath string, records [][]string)

func readFilesInDir(startPath string, r readFunc) error {
	err := filepath.WalkDir(startPath, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() && strings.HasSuffix(d.Name(), ".csv") {
			records := readCsv(path)
			r(path, records)
		}
		return err
	})
	if err != nil {
		return err
	}
	return nil
}

func readCsv(filePath string) [][]string {
	file, err := os.Open(filePath)
	logError(err, "Unable to open file "+filePath)
	defer func(file *os.File) {
		err := file.Close()
		logError(err, "Unable to close file "+filePath)
	}(file)

	reader := csv.NewReader(file)
	reader.Comma = ';'

	_, err1 := reader.Read()
	logError(err1, "Cannot read content of file "+filePath)
	records, err := reader.ReadAll()
	logError(err, "Cannot read content of file "+filePath)
	return records
}
