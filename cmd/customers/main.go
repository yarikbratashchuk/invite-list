package main

import (
	"io"
	l "log"
	"os"

	"github.com/yarikbratashchuk/invite-list/business"

	"github.com/btcsuite/btclog"
	flags "github.com/jessevdk/go-flags"
)

var log btclog.Logger

func fatalf(format string, params ...interface{}) {
	log.Criticalf(format, params...)
	os.Exit(1)
}

func setupLog(dest io.Writer, loglevel string) {
	logBackend := btclog.NewBackend(dest)
	lvl, _ := btclog.LevelFromString(loglevel)

	businessLog := logBackend.Logger("BSNS")
	log = logBackend.Logger("CUST")

	businessLog.SetLevel(lvl)
	log.SetLevel(lvl)

	business.UseLogger(businessLog)
}

func main() {
	conf, err := loadConfig()
	if err != nil {
		if flagsErr, ok := err.(*flags.Error); ok &&
			flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		}
		l.Fatalf("loading config: %v\n", err)
	}

	setupLog(os.Stderr, conf.LogLevel)

	inputf, err := os.Open(conf.Input)
	if err != nil {
		fatalf("reading %s: %v", conf.Input, err)
	}
	defer inputf.Close()

	outputf := os.Stdout
	if conf.Output != "" {
		var err error
		outputf, err = os.Create(conf.Output)
		if err != nil {
			fatalf("creating %s: %v", conf.Output, err)
		}
		defer outputf.Close()
	}

	customers, err := business.ReadCustomers(inputf)
	if err != nil {
		os.Exit(1)
	}

	invitedCustomers := business.InviteCustomers(
		customers,
		conf.Office,
		conf.MaxDistance*1000,
	)
	business.SortCustomersByID(invitedCustomers)

	business.WriteCustomers(outputf, invitedCustomers)
}
