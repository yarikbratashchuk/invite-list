package main

import (
	flags "github.com/jessevdk/go-flags"
	"github.com/yarikbratashchuk/invite-list/business"
)

type config struct {
	Office      business.Office `long:"office" description:"Office location {SF, Chicago, SurryHills, Dublin, London}"`
	MaxDistance uint            `short:"d" long:"max-invite-distance" description:"The largest distance (km) to invite customer to the office"`

	Input  string `short:"i" long:"input" description:"File with customer records"`
	Output string `short:"o" long:"output" description:"Output file"`

	LogLevel string `long:"loglevel" description:"Logging level for all subsystems {trace, debug, info, error, critical}"`
}

var defconf = config{
	Office:      business.Dublin,
	MaxDistance: 100,

	Input:  "all-customers.txt",
	Output: "",

	LogLevel: "info",
}

func loadConfig() (*config, error) {
	conf := defconf
	_, err := flags.Parse(&conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
