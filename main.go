package main

import (
	"log"
	"os/exec"
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
		File:     "assets/data.gps.json",
		Log:      "gps.log",
	},
	&devices.Bme280{
		Name:    "Bme 280",
		Device:  "/dev/i2c-1",
		Address: 0x76,
		File:    "assets/data.bme.json",
		Log:     "bme.log",
	},
}

var wg sync.WaitGroup

func main() {
	go execCommand("php", "-S", "192.168.0.14:3000")
	for _, device := range dvs {
		wg.Add(1)
		go pkg.WritingToJSON(device)
	}
	wg.Wait()
}

func execCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
