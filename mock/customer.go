package mock

import (
	"github.com/yarikbratashchuk/invite-list/business"
	"github.com/yarikbratashchuk/invite-list/geo"
)

var (
	ValidCustomerJson = `{"latitude": "52.986375", "user_id": 12, "name": "Christina McArdle", "longitude": "-6.043701"}`
	ValidCustomer     = business.Customer{
		ID:   12,
		Name: "Christina McArdle",
		Coords: geo.Coords{
			Lat:  52.986375,
			Long: -6.043701,
		},
	}
)
