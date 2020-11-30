package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	jsonFile, err := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Println("Error al abrir archivo:", err)
	}
	var i int
	for {
		jsonFile.Seek(0, 0)
		_, err = fmt.Fprintf(jsonFile, `{ "Numero": %4d }`, i)
		if err != nil {
			log.Println("Error al codificar:", err)
		}
		i++
		if i > 200 {
			i = 0
		}
		time.Sleep(1000 * time.Millisecond)
	}
}