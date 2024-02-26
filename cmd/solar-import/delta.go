package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func importPowerDelta(path string) error {

	csv := readCsv(path)

	for _, line := range csv {
		if len(line[1]) > 0 {
			unixSeconds, err := strconv.ParseInt(line[0], 10, 64)
			stringOneVoltage, err := strconv.ParseFloat(line[1], 64)
			stringOneAmp, err := strconv.ParseFloat(line[2], 64)
			stringTwoVoltage, err := strconv.ParseFloat(line[6], 64)
			stringTwoAmp, err := strconv.ParseFloat(line[7], 64)
			logError(err, "Invalid Number")

			timestamp := time.Unix(unixSeconds, 0)

			log.Println("Writing power delta data for: ", timestamp)
			writeValue("Kostal_Inverter_PV_Str1_Amperage", timestamp, stringOneAmp/1000)
			writeValue("Kostal_Inverter_PV_Str1_Voltage", timestamp, stringOneVoltage)
			writeValue("Kostal_Inverter_PV_Str2_Amperage", timestamp, stringTwoAmp/1000)
			writeValue("Kostal_Inverter_PV_Str2_Voltage", timestamp, stringTwoVoltage)

			fmt.Println("Date:", timestamp, "String 1 V:", stringOneVoltage, "String 1 A:", stringOneAmp/1000, "String 2 V:", stringTwoVoltage, "String 2 A:", stringTwoAmp/1000)
		}
	}

	return nil
}
