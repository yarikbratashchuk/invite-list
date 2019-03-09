// Package geo holds all GPS related stuff
package geo

import "math"

// Locator is used to get the location (GPS coordinates) of the object
type Locator interface {
	Location() Coords
}

// Coords holds GPS latitude and longitude values (in decimal degrees)
type Coords struct {
	Lat  float64 `json:"latitude,string"`
	Long float64 `json:"longitude,string"`
}

const earthRadius float64 = 6378100

// Distance function returns the distance (in meters) between loc1 and loc2
// through the Haversin Distance Formula for great arc distance on a
// sphere with accuracy for small distances
// See http://en.wikipedia.org/wiki/Haversine_formula
func Distance(loc1, loc2 Locator) float64 {
	l1, l2 := loc1.Location(), loc2.Location()

	l1.Lat = l1.Lat * math.Pi / 180
	l1.Long = l1.Long * math.Pi / 180
	l2.Lat = l2.Lat * math.Pi / 180
	l2.Long = l2.Long * math.Pi / 180

	d := haversin(l2.Lat-l1.Lat) +
		math.Cos(l1.Lat)*math.Cos(l2.Lat)*haversin(l2.Long-l1.Long)

	return 2 * earthRadius * math.Asin(math.Sqrt(d))
}

func haversin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}
