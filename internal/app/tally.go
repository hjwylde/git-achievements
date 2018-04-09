package app

import "encoding/json"
import "flag"
import "github.com/hjwylde/git-achievements/internal/pkg/git"
import "github.com/hjwylde/git-achievements/internal/pkg/groups"
import "github.com/hjwylde/git-achievements/internal/pkg/notes"
import "os/exec"
import "strings"

var tallyCmd = &Command{
	Run: runTallyCmd,
}

var tallyFlagSet = flag.NewFlagSet("tally", flag.ExitOnError)

func runTallyCmd(args []string) error {
	tallyFlagSet.Parse(args)

	cmd := exec.Command("git", "rev-list", "head")

	output, err := cmd.Output()
	if err != nil {
		return err
	}

	for _, sha := range strings.Fields(string(output)) {
		cmd := exec.Command("git", "notes", "--ref=achievements/progress", "remove", "--ignore-missing", sha)

		err := cmd.Run()
		if err != nil {
			return err
		}

		commit := git.Commit{
			Sha: sha,
		}

		for _, achievement := range groups.All {
			err = tally(commit, achievement)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func tally(commit git.Commit, achievement notes.Achievement) error {
	if ok, err := achievement.Match(commit); !ok || err != nil {
		return err
	}

	progress := &notes.Progress{
		Group: achievement.Group,
		Code:  achievement.Code,
	}

	b, err := json.Marshal(progress)
	if err != nil {
		return err
	}

	cmd := exec.Command("git", "notes", "--ref=achievements/progress", "append", "-m", string(b), commit.Sha)

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
