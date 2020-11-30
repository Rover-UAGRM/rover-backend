package pkg

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

// JSONAble interface
type JSONAble interface {
	Init() error
	GetFilePath() string
	GetName() string
}

// WritingToJSON function
func WritingToJSON(device JSONAble) {
	if err := device.Init(); err != nil {
		log.Println("Error al iniciar el dispositivo:", device.GetName())
	}
	jsonFile, err := os.OpenFile(device.GetFilePath(), os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Println("Error al abrir archivo:", err)
	}
	encoder := json.NewEncoder(jsonFile)
	for {
		if err = encoder.Encode(device); err != nil {
			log.Println("Error al codificar:", err)
		}
		time.Sleep(1000 * time.Millisecond)
	}
}
