package main

import (
	"log"
	filepath2 "path/filepath"
	"strconv"
	"strings"
	"time"
)

func importEnergyData(path string) error {
	err := readFilesInDir(path, handleEnergyCsvContent)
	if err != nil {
		return err
	}
	return nil
}

func handleEnergyCsvContent(filepath string, records [][]string) {
	_, file := filepath2.Split(filepath)
	before, _ := strings.CutSuffix(file, ".csv")
	split := strings.Split(before, "-")

	year, err := strconv.Atoi(split[0])
	month, err := strconv.Atoi(split[1])
	logError(err, "Invalid date string")

	for _, line := range records {
		day, err := strconv.Atoi(line[0])
		logError(err, "Invalid date string")
		energy, err := strconv.ParseFloat(line[1], 64)
		logError(err, "Invalid number")

		parsedTime := time.Date(year, time.Month(month), day, 23, 0, 0, 0, location)

		log.Println("Writing energy data for: ", parsedTime)
		writeValue("Kostal_Inverter_Yield_Day", parsedTime, energy/1000)
	}
}
