package app

var CommandMap = map[string]*Command{
	"log":            logCmd,
	"prune":          pruneCmd,
	"pull":           pullCmd,
	"push":           pushCmd,
	"tally":          tallyCmd,
	"tally-progress": tallyProgressCmd,
	"tally-unlocked": tallyUnlockedCmd,
}
