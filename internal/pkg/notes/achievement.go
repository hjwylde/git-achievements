package notes

import "github.com/hjwylde/git-achievements/internal/pkg/git"

type Achievement struct {
	Group string
	Code  string
	Name  string
	Match func(git.Commit) (bool, error)
	Count int
}

func (this Achievement) NewProgress() Progress {
	return Progress{
		Ok: true,
	}
}

func (this Achievement) NewUnlocked() Unlocked {
	return Unlocked{
		Group: this.Group,
		Code:  this.Code,
		Name:  this.Name,
	}
}

func (this Achievement) IsUnlocked(progress int) bool {
	return progress >= this.Count
}
