package app

import "encoding/json"
import "flag"
import "github.com/hjwylde/git-achievements/internal/pkg/git"
import "github.com/hjwylde/git-achievements/internal/pkg/groups"
import "github.com/hjwylde/git-achievements/internal/pkg/notes"

var tallyCmd = &Command{
	Run: runTallyCmd,
}

var tallyFlagSet = flag.NewFlagSet("tally", flag.ExitOnError)

func runTallyCmd(args []string) error {
	tallyFlagSet.Parse(args)

	err := tallyProgress()

	return err
}

func tallyProgress() error {
	revisions, err := listRevisions("head")
	if err != nil {
		return err
	}

	for _, achievement := range groups.All {
		ref := notes.AchievementProgressRef(achievement)

		for _, sha := range revisions {
			commit := git.Commit{
				Sha: sha,
			}

			err = tally(commit, achievement, ref)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func tally(commit git.Commit, achievement notes.Achievement, ref string) error {
	if ok, err := achievement.Match(commit); !ok || err != nil {
		return err
	}

	progress := achievement.NewProgress()

	b, err := json.Marshal(progress)
	if err != nil {
		return err
	}

	err = addNote(string(b), commit.Sha, ref)

	return err
}
