package app

import "flag"
import "github.com/hjwylde/git-achievements/internal/pkg/gexec"
import "github.com/hjwylde/git-achievements/internal/pkg/notes"

var pruneCmd = &Command{
	FlagSet: flag.NewFlagSet("prune", flag.ExitOnError),
	Run: func(flagSet *flag.FlagSet) error {
		return prune()
	},
}

func prune() error {
	ref := notes.ProgressRef

	err := gexec.PruneNotes(ref)

	return err
}
