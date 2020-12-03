package main

import (
	"time"

	"./pkg"
	"./pkg/devices"
)

type devs []pkg.Device

var dvs = devs{
	&devices.Gps{
		Name:     "GPS Neo 6m",
		Device:   "/dev/ttyS0",
		Baudrate: 115200,
		File:     "data.gps.json",
	},
}

func main() {
	for _, device := range dvs {
		go pkg.WritingToJSON(device)
	}
	time.Sleep(5000 * time.Millisecond)
}
