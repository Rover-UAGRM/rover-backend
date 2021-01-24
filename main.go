package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"sync"

	"./pkg"
	"./pkg/devices"
)

type deviceListType []pkg.Device

var deviceList = deviceListType{
	&devices.Gps{
		Name:     "GPS Neo 6m",
		Device:   "/dev/ttyS0",
		Baudrate: 115200,
		File:     "assets/data.gps.json",
		Log:      "gps.log",
	},
	&devices.Bme280{
		Name:         "Bme 280",
		Device:       "/dev/i2c-1",
		SamplingTime: 250,
		Address:      0x76,
		File:         "assets/data.bme.json",
		Log:          "bme.log",
	},
}

var wg sync.WaitGroup

func main() {

	host := flag.String("host", "127.0.0.1", "Host address")
	port := flag.String("port", "3000", "Usage Port")
	// mode := flag.String("mode", "tx", "Rover Modality")
	flag.Parse()

	go execCommand("php", "-S", fmt.Sprintf("%s:%s", *host, *port))
	for _, device := range deviceList {
		wg.Add(1)
		go pkg.WritingToJSON(device, 1000)
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
