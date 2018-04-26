package app

import "encoding/json"
import "flag"
import "github.com/hjwylde/git-achievements/internal/pkg/gexec"
import "github.com/hjwylde/git-achievements/internal/pkg/git"
import "github.com/hjwylde/git-achievements/internal/pkg/groups"
import "github.com/hjwylde/git-achievements/internal/pkg/notes"

var tallyProgressCmd = &Command{
	FlagSet: flag.NewFlagSet("tally-progress", flag.ExitOnError),
	Run: func(flagSet *flag.FlagSet) error {
		return tallyProgress()
	},
}

func tallyProgress() error {
	revisions, err := gexec.ListRevisions("head")
	if err != nil {
		return err
	}

	for _, achievement := range groups.All {
		ref := notes.AchievementProgressRef(achievement)

		for _, sha := range revisions {
			commit := git.Commit{
				Sha: sha,
			}

			ok, err := achievement.Match(commit)
			if err != nil {
				return err
			}
			if !ok {
				continue
			}

			progress := achievement.NewProgress()

			b, err := json.Marshal(progress)
			if err != nil {
				return err
			}

			err = gexec.AddNote(string(b), commit.Sha, ref)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
