package app

import "flag"
import "github.com/hjwylde/git-achievements/internal/pkg/notes"

var pullCmd = &Command{
	Run: runPullCmd,
}

var pullFlagSet = flag.NewFlagSet("pull", flag.ExitOnError)

func runPullCmd(args []string) error {
	pullFlagSet.Parse(args)

	remote, err := getUpstreamRemote()
	if err != nil {
		return err
	}

	ref := notes.BaseRef

	err = fetch(remote, "refs/notes/"+ref+"/*")
	if err != nil {
		return err
	}

	err = mergeNotes(remote, ref+"/*")

	return err
}
