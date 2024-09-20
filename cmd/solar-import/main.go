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
	deltaWithEnergy := flag.Bool("delta-energy", false, "Enabled Energy calculation for delta Power import")
	deltaPath := flag.String("delta-path", ".", "Path to Power Delta File")
	dryRun := flag.Bool("dry-run", false, "Enable simulation without writing to the database")

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
		err := importPowerData(*powerPath, *dryRun)
		panicOnError(err)
	}

	if *energy {
		err := importEnergyData(*energyPath, *dryRun)
		panicOnError(err)
	}

	if *delta {
		err := importPowerDelta(*deltaPath, *deltaWithEnergy, *dryRun)
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
