// Package geo holds all GPS related stuff
package geo

import (
	"fmt"
	"math"
)

// Locator is used to get the location (GPS coordinates) of the object
type Locator interface {
	Location() Coords
}

// Coords holds GPS latitude and longitude values (in decimal degrees)
type Coords struct {
	Lat  float64 `json:"latitude,string"`
	Long float64 `json:"longitude,string"`
}

// earthRadius is a mean earth radius in meters (by IUGG)
const earthRadius float64 = 6371008.8

// Distance function returns the distance (in meters) between loc1 and loc2
//
//
// See http://en.wikipedia.org/wiki/Haversine_formula for details on formula.
func Distance(loc1, loc2 Locator) (float64, error) {
	l1, l2 := loc1.Location(), loc2.Location()

	if !ValidLat(l1.Lat) || !ValidLong(l1.Long) {
		return 0, fmt.Errorf("invalid coords %v", l1)
	}
	if !ValidLat(l2.Lat) || !ValidLong(l2.Long) {
		return 0, fmt.Errorf("invalid coords %v", l2)
	}

	l1.Lat = l1.Lat * math.Pi / 180
	l1.Long = l1.Long * math.Pi / 180
	l2.Lat = l2.Lat * math.Pi / 180
	l2.Long = l2.Long * math.Pi / 180

	h := haversin(l2.Lat-l1.Lat) +
		math.Cos(l1.Lat)*math.Cos(l2.Lat)*haversin(l2.Long-l1.Long)

	dist := 2 * earthRadius * math.Asin(math.Sqrt(h))

	return dist, nil
}

func haversin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

// ValidLat returns true if lat is a valid decimal degree latitude value
func ValidLat(lat float64) bool {
	if lat < float64(-90) || lat > float64(90) {
		return false
	}
	return true
}

// ValidLong returns true if lat is a valid decimal degree latitude value
func ValidLong(long float64) bool {
	if long < float64(-180) || long > float64(180) {
		return false
	}
	return true
}
