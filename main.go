package main

import (
	"./pkg"
	"./pkg/devices"
)

func main() {

	gps := devices.Gps{
		Name:     "GPS Neo 6m",
		Device:   "/uart",
		Baudrate: 115200,
		File:     "data.gps.json",
	}

	pkg.WritingToJSON(&gps)
}
