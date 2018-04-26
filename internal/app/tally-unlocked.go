package app

import "encoding/json"
import "flag"
import "github.com/hjwylde/git-achievements/internal/pkg/gexec"
import "github.com/hjwylde/git-achievements/internal/pkg/groups"
import "github.com/hjwylde/git-achievements/internal/pkg/notes"

var tallyUnlockedCmd = &Command{
	FlagSet: flag.NewFlagSet("tally-unlocked", flag.ExitOnError),
	Run: func(flagSet *flag.FlagSet) error {
		return tallyUnlocked()
	},
}

func tallyUnlocked() error {
	for _, achievement := range groups.All {
		progressRef := notes.AchievementProgressRef(achievement)

		output, err := gexec.ListNotes(progressRef)
		if err != nil {
			return err
		}

		progress := len(output)

		if ok := achievement.IsUnlocked(progress); !ok {
			continue
		}

		unlocked := achievement.NewUnlocked()

		b, err := json.Marshal(unlocked)
		if err != nil {
			return err
		}

		unlockedRef := notes.AchievementUnlockedRef(achievement)

		err = gexec.AddNote(string(b), "head", unlockedRef)
		if err != nil {
			return err
		}
	}

	return nil
}
