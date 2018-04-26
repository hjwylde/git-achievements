package git

import "time"

type Commit struct {
	Sha        string
	AuthorDate time.Time
}
