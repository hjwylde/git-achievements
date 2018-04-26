package app

import "flag"
import "github.com/hjwylde/git-achievements/internal/pkg/notes"

var pushCmd = &Command{
	Run: runPushCmd,
}

var pushFlagSet = flag.NewFlagSet("push", flag.ExitOnError)

func runPushCmd(args []string) error {
	pushFlagSet.Parse(args)

	remote, err := getPushRemote()
	if err != nil {
		return err
	}

	ref := notes.BaseRef

	err = push(remote, "refs/notes/"+ref+"/*")

	return err
}
