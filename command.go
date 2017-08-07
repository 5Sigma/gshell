package gshell

type (
	Command struct {
		Name        string
		Description string
		Call        func(*Shell, []string)
	}
)
