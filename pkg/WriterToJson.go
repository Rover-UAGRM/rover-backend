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
func WritingToJSON(device Device) {
	if err := device.Init(); err != nil {
		log.Println("Error al iniciar el dispositivo:", device.GetName())
	}
	jsonFile, err := os.OpenFile(device.GetFilePath(), os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		device.LogPrintln("Error al abrir archivo:", err)
	}
	encoder := json.NewEncoder(jsonFile)
	for {
		deleteFile(jsonFile)
		if err = encoder.Encode(device); err != nil {
			device.LogPrintln("Error al codificar:", err)
		}
		time.Sleep(1000 * time.Millisecond)
	}
}

func deleteFile(file *os.File) {
	file.Truncate(0)
	file.Seek(0, 0)
}
