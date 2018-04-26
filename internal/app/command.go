package app

import "flag"

type Command struct {
	FlagSet *flag.FlagSet
	Run     func(*flag.FlagSet) error
}
