package app

type Command struct {
	Run func([]string) error
}
