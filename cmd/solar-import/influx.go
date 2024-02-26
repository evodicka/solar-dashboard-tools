package main

import (
	"context"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"time"
)

var client influxdb2.Client

func connect(url string, username string, password string) {
	client = influxdb2.NewClient(url, username+":"+password)
}

func writeValue(measurement string, time time.Time, value interface{}) error {
	p := influxdb2.NewPointWithMeasurement(measurement).
		AddTag("item", measurement).
		AddField("value", value).
		SetTime(time)

	writeAPI := client.WriteAPIBlocking("my-org", "openhab")
	err := writeAPI.WritePoint(context.Background(), p)
	return err
}

func closeConnection() {
	if client != nil {
		client.Close()
	}
}
