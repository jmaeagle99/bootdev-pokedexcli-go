package command

type CliCommand struct {
	Name        string
	Description string
	Callback    func() error
}