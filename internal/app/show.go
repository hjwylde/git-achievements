package app

import "flag"
import "fmt"
import "github.com/hjwylde/git-achievements/internal/pkg/notes"
import "os/exec"
import "strings"

var showCmd = &Command{
	Run: runShowCmd,
}

var showFlagSet = flag.NewFlagSet("show", flag.ExitOnError)

func runShowCmd(args []string) error {
	showFlagSet.Parse(args)

	cmd := exec.Command("git", "notes", "--ref="+notes.ProgressRef, "show")

	b, err := cmd.Output()
    if err != nil {
        return err
    }

    out := strings.TrimSpace(string(b))

    fmt.Printf("%s\n", out)

	return nil
}
