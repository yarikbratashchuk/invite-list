package business

import (
	"sync"

	"github.com/yarikbratashchuk/invite-list/geo"
)

type Office string

const (
	SF         Office = "SF"
	Chicago    Office = "Chicago"
	SurryHills Office = "SurryHills"
	Dublin     Office = "Dublin"
	London     Office = "London"
)

var (
	// officeLocations holds GPS coordinates of the offices.
	// Office locations are taken from https://craft.co/intercom
	// and GPS coordinates are produced by converting locations using
	// https://www.latlong.net/convert-address-to-lat-long.html
	officeLocations = map[Office]geo.Coords{
		SF:         geo.Coords{37.788880, -122.400280},
		Chicago:    geo.Coords{41.886750, -87.659980},
		SurryHills: geo.Coords{-33.881400, 151.214400},
		Dublin:     geo.Coords{53.339428, -6.257664},
		London:     geo.Coords{51.489520, -0.088370},
	}

	officeLocationsMu sync.RWMutex
)

// Location implements geo.Locator
func (o Office) Location() geo.Coords {
	officeLocationsMu.RLock()
	c := officeLocations[o]
	officeLocationsMu.RUnlock()

	return c
}
