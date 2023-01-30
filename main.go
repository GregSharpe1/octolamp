package main

import (
	"fmt"
	"os"

	"github.com/GregSharpe1/octolamp/internal/wled"
	"github.com/sirupsen/logrus"
)

func getEnv(name string) string {
	value, ok := os.LookupEnv(name)
	if !ok {
		logrus.Fatalf("Required environment '%s' does not have a value.", name)
	}

	return value
}

func Default(name string, defaultValue string) string {
	value, ok := os.LookupEnv(name)
	if !ok {
		return defaultValue
	}
	return value
}

func main() {
	fmt.Printf("Starting OctoLamp...")

	var ipaddress string = getEnv("WLED_IP_ADDRESS")
	var schema string = Default("WLED_SCHEMA", "http")

	deviceConfig := wled.DeviceConfig{
		IPAddress: ipaddress,
		Schema:    schema,
	}

	deviceConfig.ChangeColour(deviceConfig, "red")

}
