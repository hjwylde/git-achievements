package notes

import "github.com/hjwylde/git-achievements/internal/pkg/git"

type Achievement struct {
	Group string
	Code  string
	Name  string
	Match func(git.Commit) (bool, error)
}
