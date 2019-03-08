package main

import (
	"fmt"
	"os"

	"github.com/yarikbratashchuk/invite-list/customers"

	log "github.com/btcsuite/btclog"
	flags "github.com/jessevdk/go-flags"
)

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
