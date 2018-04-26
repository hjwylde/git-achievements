package groups

import "github.com/hjwylde/git-achievements/internal/pkg/git"
import "github.com/hjwylde/git-achievements/internal/pkg/notes"

var Git = [...]notes.Achievement{
	firstCommit,
}

var firstCommit = notes.Achievement{
	Group: "git",
	Code:  "first-commit",
	Name:  "First commit!",
	Match: func(_ git.Commit) (bool, error) {
		return true, nil
	},
	Count: 1,
}
