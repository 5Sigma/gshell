package gshell

import (
	"github.com/chzyer/readline"
)

func setupAutoCompleter(sh *Shell) *readline.PrefixCompleter {
	commandStrings := []readline.PrefixCompleterInterface{}
	for _, cmd := range sh.Commands {
		commandStrings = append(commandStrings, readline.PcItem(cmd.Name))
	}
	return readline.NewPrefixCompleter(commandStrings...)
}
