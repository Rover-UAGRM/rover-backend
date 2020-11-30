package pkg

import (
	"fmt"
	"log"
	"os"
	"time"
)

// JSONAble interface
type JSONAble interface {
	Init() error
}

// WritingToJSON function
func WritingToJSON(device JSONAble, path string) {

	device.Init()

	jsonFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Println("Error al abrir archivo:", err)
	}
	for {
		jsonFile.Seek(0, 0)
		_, err = fmt.Fprintln(jsonFile, `device.String()`)
		if err != nil {
			log.Println("Error al codificar:", err)
		}
		time.Sleep(1000 * time.Millisecond)
	}
}
