package devices

import (
	"errors"
	"fmt"
	"log"
	"math"
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

	Status, ns, ew    byte
	latitud, longitud float64
	Latitud, Longitud float64
}

var intValue int
var floatValue float64

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
		gps.logger = log.New(logFile, "Logger: ", log.Llongfile)
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
			gps.Latitud = convertDegMinToDecDeg(gps.latitud)
			gps.Longitud = convertDegMinToDecDeg(gps.longitud)

			if gps.ns == 'S' {
				gps.Latitud *= -1
			}

			if gps.ew == 'W' {
				gps.Longitud *= -1
			}
		}

	}
}

func convertDegMinToDecDeg(degMin float64) float64 {
	min := 0.0
	decDeg := 0.0
	min = math.Mod(degMin, 100)
	degMinInt := int(degMin) / 100
	decDeg = float64(degMinInt) + (min / 60)
	return roundTo(decDeg, 6)
}

func roundTo(n float64, decimals uint32) float64 {
	return math.Round(n*math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
}
