package devices

import (
	"errors"
	"log"
	"os"
	"time"

	"golang.org/x/exp/io/i2c"

	"github.com/quhar/bme280"
)

//Bme280 : Estructura que contiene datos del sensor de presion
type Bme280 struct {
	Name      string
	Address   int    `json:"-"`
	Device    string `json:"-"`
	File      string `json:"-"`
	Log       string `json:"-"`
	logger    *log.Logger
	devReader *bme280.BME280

	Temperature, Pressure, Humidity float64
}

//GetName function
func (bme *Bme280) GetName() string {
	return bme.Name
}

//GetFilePath function
func (bme *Bme280) GetFilePath() string {
	return bme.File
}

//InitLogger function
func (bme *Bme280) initLogger() error {
	if bme.Log != "" {
		logFile, err := os.OpenFile(bme.Log, os.O_RDWR|os.O_CREATE, os.ModePerm)
		if err != nil {
			return err
		}
		bme.logger = log.New(logFile, "Log: ", log.LstdFlags)
		return err
	} else {
		return errors.New("Not found log path")
	}
}

//LogPrintln function
func (bme *Bme280) LogPrintln(v ...interface{}) {
	bme.logger.Println(v...)
}

//Init : Inicializa
func (bme *Bme280) Init() error {
	err := bme.initLogger()
	if err != nil {
		return err
	}
	i2c, errInitI2c := i2c.Open(&i2c.Devfs{Dev: bme.Device}, bme.Address)
	if errInitI2c != nil {
		return errInitI2c
	}

	bme.devReader = bme280.New(i2c)
	errInit := bme.devReader.Init()
	if errInit != nil {
		return err
	}

	go bme.reading()
	return err
}

//Reading : Rutina para leer constantemente los datos del bme
func (bme *Bme280) reading() {
	for {
		t, p, h, err := bme.devReader.EnvData()
		if err != nil {
			bme.LogPrintln("Error Reading: ", err)
		}

		bme.Temperature = roundTo(t, 0)
		bme.Pressure = roundTo(p, 2)
		bme.Humidity = roundTo(h, 0)

		time.Sleep(250 * time.Millisecond)
	}
}
