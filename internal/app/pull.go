package app

import "flag"
import "github.com/hjwylde/git-achievements/internal/pkg/gexec"
import "github.com/hjwylde/git-achievements/internal/pkg/notes"

var pullCmd = &Command{
	FlagSet: flag.NewFlagSet("pull", flag.ExitOnError),
	Run: func(flagSet *flag.FlagSet) error {
		return pull()
	},
}

func pull() error {
	remote, err := gexec.GetUpstreamRemote()
	if err != nil {
		return err
	}

	ref := notes.BaseRef

	err = gexec.Fetch(remote, "refs/notes/"+ref+"/*")
	if err != nil {
		return err
	}

	err = gexec.MergeNotes(remote, ref+"/*")

	return err
}
