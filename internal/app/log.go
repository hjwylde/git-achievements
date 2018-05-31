package app

import "flag"

var logCmd = &Command{
	FlagSet: flag.NewFlagSet("log", flag.ExitOnError),
	Run: func(flagSet *flag.FlagSet) error {
		return log()
	},
}

func log() error {
		unlockedRef := notes.UnlockedRef+"/*"

		unlockedByCommit, err := gexec.GetUnlocked(unlockedRef)
		if err != nil {
			return err
		}

		commits := make([]git.Commit, 0, len(unlockedByCommit))
		for commit := range unlockedByCommit {
			commits = append(commits, commit)
		}

		sort.Slice(commits, func(i, j int) bool {
			return commits[i].AuthorDate.Before(commits[j].AuthorDate)
		})

		for _, commit := range commits {
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

	return nil
}
