package app

var CommandMap = map[string]*Command{
	"prune": pruneCmd,
	"pull":  pullCmd,
	"push":  pushCmd,
	"show":  showCmd,
	"tally": tallyCmd,
}
