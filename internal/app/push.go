package app

import "flag"
import "github.com/hjwylde/git-achievements/internal/pkg/notes"
import "os/exec"

var pushCmd = &Command{
	Run: runPushCmd,
}

// TODO (hjw): allow customising the remote
// From git pull docs: default values for <repository> and <branch> are read from the "remote" and "merge" configuration for the current branch as set by git-branch(1) --track.
var pushFlagSet = flag.NewFlagSet("push", flag.ExitOnError)

func runPushCmd(args []string) error {
	pushFlagSet.Parse(args)

	cmd := exec.Command("git", "push", "origin", "refs/notes/"+notes.BaseRef+"/*")

	err := cmd.Run()

	return err
}
