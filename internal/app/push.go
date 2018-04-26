package app

import "flag"
import "github.com/hjwylde/git-achievements/internal/pkg/gexec"
import "github.com/hjwylde/git-achievements/internal/pkg/notes"

var pushCmd = &Command{
	FlagSet: flag.NewFlagSet("push", flag.ExitOnError),
	Run: func(flagSet *flag.FlagSet) error {
		return push()
	},
}

func push() error {
	remote, err := gexec.GetPushRemote()
	if err != nil {
		return err
	}

	ref := notes.BaseRef

	err = gexec.Push(remote, "refs/notes/"+ref+"/*")

	return err
}
