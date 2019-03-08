package main

import (
	"fmt"
	"os"

	"github.com/yarikbratashchuk/invite-list/customers"

	log "github.com/btcsuite/btclog"
	flags "github.com/jessevdk/go-flags"
)

type config struct {
	Office      customers.Office `long:"office" description:"Office location {SF, Chicago, SurryHills, Dublin, London}"`
	MaxDistance int              `short:"d" long:"max-invite-distance" description:"The largest distance (km) to invite customer to the office"`

	File string `short:"f" long:"from-file" description:"File with customer records"`
	URL  string `short:"u" long:"from-url" description:"Url to load customer records from"`

	Output string `short:"o" long:"output" description:"Output file"`

	LogLevel string `long:"loglevel" description:"Logging level for all subsystems {trace, debug, info, error, critical}"`
}

var defconf = config{
	Office:      customers.Dublin,
	MaxDistance: 100,
	File:        "customers.txt",
	Output:      "invite-customers.txt",

	LogLevel: "debug",
}

func loadConfig() (*config, error) {
	conf := defconf
	_, err := flags.Parse(&conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}

func main() {
	conf, err := loadConfig()
	if err != nil {
		if flagsErr, ok := err.(*flags.Error); ok &&
			flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		}
		fmt.Printf("loading config: %v\n", err)
		os.Exit(1)
	}

	// setup logging
	logBackend := log.NewBackend(os.Stderr)
	lvl, _ := log.LevelFromString(conf.LogLevel)

	customersLog := logBackend.Logger("CUST")
	customersLog.SetLevel(lvl)
	customers.UseLogger(customersLog)

	fmt.Println(conf)
}
