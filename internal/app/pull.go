package app

import "flag"
import "github.com/hjwylde/git-achievements/internal/pkg/notes"
import "os/exec"

var pullCmd = &Command{
	Run: runPullCmd,
}

// TODO (hjw): allow customising the remote
// From git pull docs: default values for <repository> and <branch> are read from the "remote" and "merge" configuration for the current branch as set by git-branch(1) --track.
var pullFlagSet = flag.NewFlagSet("pull", flag.ExitOnError)

func runPullCmd(args []string) error {
	pullFlagSet.Parse(args)

	cmd := exec.Command("git", "fetch", "origin", "refs/notes/"+notes.BaseRef+"/*")

	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("git", "notes", "merge", "-s cat_sort_uniq", "origin/"+notes.BaseRef+"/*")

    err := cmd.Run()

	return err
}
