package app

import "encoding/json"
import "flag"
import "github.com/hjwylde/git-achievements/internal/pkg/gexec"
import "github.com/hjwylde/git-achievements/internal/pkg/git"
import "github.com/hjwylde/git-achievements/internal/pkg/groups"
import "github.com/hjwylde/git-achievements/internal/pkg/notes"
import "sort"

var tallyUnlockedCmd = &Command{
	FlagSet: flag.NewFlagSet("tally-unlocked", flag.ExitOnError),
	Run: func(flagSet *flag.FlagSet) error {
		return tallyUnlocked()
	},
}

func tallyUnlocked() error {
	for _, achievement := range groups.All {
		progressRef := notes.AchievementProgressRef(achievement)

		progressByCommit, err := gexec.GetProgress(progressRef)
		if err != nil {
			return err
		}

		commits := make([]git.Commit, 0, len(progressByCommit))
		for commit := range progressByCommit {
			commits = append(commits, commit)
		}

		sort.Slice(commits, func(i, j int) bool {
			return commits[i].AuthorDate.Before(commits[j].AuthorDate)
		})

		for i, commit := range commits {
			if !achievement.IsUnlocked(i + 1) {
				continue
			}

			unlocked := achievement.NewUnlocked()

			b, err := json.Marshal(unlocked)
			if err != nil {
				return err
			}

			unlockedRef := notes.AchievementUnlockedRef(achievement)

			if err = gexec.AddNote(string(b), commit.Sha, unlockedRef); err != nil {
				return err
			}

			break
		}
	}

	return nil
}
