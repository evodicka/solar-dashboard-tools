package main

import (
	"log"
	"strconv"
	"time"
)

func importPowerData(path string) error {

	err := readFilesInDir(path, handlePowerCsvContent)
	if err != nil {
		return err
	}
	return nil

}

func handlePowerCsvContent(_ string, records [][]string) {
	for _, line := range records {
		timeString := line[0]
		stringOnePower, err := strconv.ParseFloat(line[4], 64)
		stringTwoPower, err := strconv.ParseFloat(line[5], 64)
		logError(err, "Invalid Numbers")

		parsedTime, err := time.ParseInLocation(time.DateTime, timeString, location)
		logError(err, "Invalid date string")

		log.Println("Writing power data for: ", parsedTime)
		writeValue("Kostal_Inverter_PV_Str1_Amperage", parsedTime, stringOnePower)
		writeValue("Kostal_Inverter_PV_Str1_Voltage", parsedTime, 1.0)
		writeValue("Kostal_Inverter_PV_Str2_Amperage", parsedTime, stringTwoPower)
		writeValue("Kostal_Inverter_PV_Str2_Voltage", parsedTime, 1.0)
	}
}
