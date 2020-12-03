package main

import (
	"sync"

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

var wg sync.WaitGroup

func main() {
	for _, device := range dvs {
		wg.Add(1)
		go pkg.WritingToJSON(device)
	}
	wg.Wait()
	// time.Sleep(5000 * time.Millisecond)
}
