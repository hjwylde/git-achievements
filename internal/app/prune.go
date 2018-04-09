package app

import "flag"
import "github.com/hjwylde/git-achievements/internal/pkg/notes"
import "os/exec"

var pruneCmd = &Command{
	Run: runPruneCmd,
}

var pruneFlagSet = flag.NewFlagSet("prune", flag.ExitOnError)

func runPruneCmd(args []string) error {
	pruneFlagSet.Parse(args)

	cmd := exec.Command("git", "notes", "--ref="+notes.ProgressRef, "prune")

	err := cmd.Run()

	return err
}
