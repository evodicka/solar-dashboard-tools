package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

var location, _ = time.LoadLocation("Europe/Berlin")

func main() {
	power := flag.Bool("power", false, "Import Power Data")
	energy := flag.Bool("energy", false, "Import Energy Data")
	powerPath := flag.String("power-path", ".", "Path to Power Data")
	energyPath := flag.String("energy-path", ".", "Path to Energy Data")
	delta := flag.Bool("delta", false, "Import Power Delta")
	deltaPath := flag.String("delta-path", ".", "Path to Power Delta File")

	databaseUrl := flag.String("d", "http://localhost:8086", "InfluxDB URL")
	databaseUser := flag.String("u", "", "Database User")
	databasePassword := flag.String("p", "", "Database Password")

	flag.Parse()

	fmt.Println("Power Import enabled:", *power, "Using Path:", *powerPath)
	fmt.Println("Energy Import enabled:", *energy, "Using Path", *energyPath)
	fmt.Println("Delta Import enabled:", *delta, "Using Path", *deltaPath)

	connect(*databaseUrl, *databaseUser, *databasePassword)
	defer closeConnection()

	if *power {
		err := importPowerData(*powerPath)
		panicOnError(err)
	}

	if *energy {
		err := importEnergyData(*energyPath)
		panicOnError(err)
	}

	if *delta {
		err := importPowerDelta(*deltaPath)
		panicOnError(err)
	}
}

func logError(err error, message string) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
