package app

var CommandMap = map[string]*Command{
	"prune": pruneCmd,
	"pull":  pullCmd,
	"push":  pushCmd,
	"tally": tallyCmd,
}
