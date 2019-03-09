package mock

import "github.com/yarikbratashchuk/invite-list/geo"

type Locator geo.Coords

func (l Locator) Location() geo.Coords { return geo.Coords(l) }

var (
	// valid
	Locator1 = Locator{0, 0}
	Locator2 = Locator{1, 1}
	Locator3 = Locator{-1, -1}

	// invalid
	LocatorInv1 = Locator{95, 0}
	LocatorInv2 = Locator{90, 181}
)
