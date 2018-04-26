package app

import "flag"

var tallyCmd = &Command{
	FlagSet: flag.NewFlagSet("tally", flag.ExitOnError),
	Run: func(flagSet *flag.FlagSet) error {
		return tally()
	},
}

func tally() error {
	if err := tallyProgress(); err != nil {
		return err
	}

	err = tallyUnlocked()

	return err
}
