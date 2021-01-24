package pkg

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

// Device interface
type Device interface {
	Init() error
	GetFilePath() string
	GetName() string
	LogPrintln(v ...interface{})
}

// WritingToJSON function
func WritingToJSON(device Device, duration time.Duration) {
	if err := device.Init(); err != nil {
		log.Println("Error al iniciar el dispositivo:", device.GetName())
	}
	jsonFile, err := os.OpenFile(device.GetFilePath(), os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		device.LogPrintln("Error al abrir archivo:", err)
	}
	encoder := json.NewEncoder(jsonFile)
	for {
		cleanUpFile(jsonFile)
		if err = encoder.Encode(device); err != nil {
			device.LogPrintln("Error al codificar:", err)
		}
		time.Sleep(duration * time.Millisecond)
	}
}

func cleanUpFile(file *os.File) {
	file.Truncate(0)
	file.Seek(0, 0)
}
