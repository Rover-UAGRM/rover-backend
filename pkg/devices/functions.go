package devices

import "math"

func roundTo(n float64, decimals uint32) float64 {
	return math.Round(n*math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
}

func convertDegMinToDecDeg(degMin float64) float64 {
	min := 0.0
	decDeg := 0.0
	min = math.Mod(degMin, 100)
	degMinInt := int(degMin) / 100
	decDeg = float64(degMinInt) + (min / 60)
	return roundTo(decDeg, 6)
}

func readAltitude(seaLevel, pressure float64) float64 {
	return roundTo(44330.0*(1.0-math.Pow(pressure/seaLevel, 0.1903)), 0)
}
