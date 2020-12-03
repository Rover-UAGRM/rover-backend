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
}

// WritingToJSON function
func WritingToJSON(device Device) {
	if err := device.Init(); err != nil {
		log.Println("Error al iniciar el dispositivo:", device.GetName())
	}
	jsonFile, err := os.OpenFile(device.GetFilePath(), os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Println("Error al abrir archivo:", err)
	}
	encoder := json.NewEncoder(jsonFile)
	for {
		jsonFile.Truncate(0)
		jsonFile.Seek(0, 0)
		if err = encoder.Encode(device); err != nil {
			log.Println("Error al codificar:", err)
		}
		time.Sleep(1000 * time.Millisecond)
	}
}
