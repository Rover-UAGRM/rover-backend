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
}

var wg sync.WaitGroup

func main() {
	go execCommand("http-server")
	for _, device := range dvs {
		wg.Add(1)
		go pkg.WritingToJSON(device)
	}
	wg.Wait()
}

func execCommand(cmdString string) {
	cmd := exec.Command(cmdString)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
