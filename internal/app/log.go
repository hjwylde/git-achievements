package app

import "flag"

var logCmd = &Command{
	FlagSet: flag.NewFlagSet("log", flag.ExitOnError),
	Run: func(flagSet *flag.FlagSet) error {
		return log()
	},
}

func log() error {
    return nil
}
