package main

import (
	"./pkg"
	"./pkg/devices"
)

func main() {

	gps := devices.Gps{
		Device:   "/uart",
		Baudrate: 115200,
	}

	pkg.WritingToJSON(&gps, "data.gps.json")
}
