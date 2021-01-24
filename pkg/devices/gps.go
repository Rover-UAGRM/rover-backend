package devices

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tarm/serial"
)

//Gps : Estructura que contiene datos del gps
type Gps struct {
	Name     string
	Device   string `json:"-"`
	Baudrate int    `json:"-"`
	File     string `json:"-"`
	Log      string `json:"-"`
	logger   *log.Logger
	port     *serial.Port

	Hour, Min, Sec   int
	Day, Month, Year int

	Status, ns, ew      byte
	latitud, longitud   float64
	Latitude, Longitude float64

	Speed float64

	NroSats int

	Altitude float64
}

var intValue int
var floatValue float64
var charValue byte

//GetName function
func (gps *Gps) GetName() string {
	return gps.Name
}

//GetFilePath function
func (gps *Gps) GetFilePath() string {
	return gps.File
}

//InitLogger function
func (gps *Gps) initLogger() error {
	if gps.Log != "" {
		logFile, err := os.OpenFile(gps.Log, os.O_RDWR|os.O_CREATE, os.ModePerm)
		if err != nil {
			return err
		}
		gps.logger = log.New(logFile, "Log: ", log.LstdFlags)
		return err
	} else {
		return errors.New("Not found log path")
	}
}

//LogPrintln function
func (gps *Gps) LogPrintln(v ...interface{}) {
	gps.logger.Println(v...)
}

//Init : Inicializa
func (gps *Gps) Init() error {
	err := gps.initLogger()
	if err != nil {
		return err
	}
	c := &serial.Config{Name: gps.Device, Baud: gps.Baudrate}
	gps.port, err = serial.OpenPort(c)
	if err != nil {
		return err
	}
	go gps.reading()
	return err
}

//Reading : Rutina para leer constantemente los datos del Gps
func (gps *Gps) reading() {
	var nmeaString string
	for {
		nmeaString = ""
		_, err := fmt.Fscanln(gps.port, &nmeaString)
		if err != nil {
			gps.LogPrintln("Error Escaneo:", err)
			continue
		}

		if strings.Contains(nmeaString, "$GPRMC") {
			_, err = fmt.Sscanf(nmeaString, "$GPRMC,%2d%2d%2d.%2d,%c,%f,%c,%f,%c,%f,,%2d%2d%2d,",
				&gps.Hour, &gps.Min, &gps.Sec,
				&intValue,
				&gps.Status,
				&gps.latitud, &gps.ns,
				&gps.longitud, &gps.ew,
				&floatValue,
				&gps.Day, &gps.Month, &gps.Year)
			if err != nil {
				gps.LogPrintln("Error GPRMC:", err)
				continue
			}
			gps.Latitude = convertDegMinToDecDeg(gps.latitud)
			gps.Longitude = convertDegMinToDecDeg(gps.longitud)

			if gps.ns == 'S' {
				gps.Latitude *= -1
			}

			if gps.ew == 'W' {
				gps.Longitude *= -1
			}
		}

		if strings.Contains(nmeaString, "$GPVTG") {
			_, err = fmt.Sscanf(nmeaString, "$GPVTG,,%c,,%c,%f,%c,%f",
				&charValue, &charValue, &floatValue, &charValue,
				&gps.Speed)
			if err != nil {
				gps.LogPrintln("Error GPVTG:", err)
				continue
			}
			gps.Speed = roundTo(gps.Speed, 0)
		}

		if strings.Contains(nmeaString, "$GPGSV") {
			_, err = fmt.Sscanf(nmeaString, "$GPGSV,%d,%d,%d",
				&intValue, &intValue,
				&gps.NroSats)
			if err != nil {
				gps.LogPrintln("Error GPGSV:", err)
				continue
			}
		}

		if strings.Contains(nmeaString, "$GPGGA") {
			_, err = fmt.Sscanf(nmeaString, "$GPGGA,%f,%f,%c,%f,%c,%c,%d,%f,%f",
				&floatValue, &floatValue, &charValue, &floatValue, &charValue, &charValue, &intValue, &floatValue,
				&gps.Altitude)
			if err != nil {
				gps.LogPrintln("Error GPGGA:", err)
				continue
			}
			gps.Altitude = roundTo(gps.Altitude, 0)
		}
	}
}
