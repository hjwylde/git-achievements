package app

import "flag"
import "github.com/hjwylde/git-achievements/internal/pkg/notes"

var pruneCmd = &Command{
	Run: runPruneCmd,
}

var pruneFlagSet = flag.NewFlagSet("prune", flag.ExitOnError)

func runPruneCmd(args []string) error {
	pruneFlagSet.Parse(args)

	ref := notes.ProgressRef

	err := pruneNotes(ref)

	return err
}
