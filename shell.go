package gshell

import (
	"fmt"
	"github.com/chzyer/readline"
	"strings"
)

type (
	Shell struct {
		Commands []*Command
		Prompt   string
	}
)

func New() *Shell {
	sh := &Shell{
		Prompt:   "> ",
		Commands: []*Command{},
	}
	sh.AddCommand(&Command{
		Name:        "help",
		Description: "Show help text",
		Call: func(sh *Shell, args []string) {
			sh.ShowHelp()
		},
	})
	return sh
}

func (shell *Shell) AddCommand(cmd *Command) {
	shell.Commands = append(shell.Commands, cmd)
}

func (shell *Shell) Start() error {

	config := &readline.Config{
		Prompt:            shell.Prompt,
		AutoComplete:      setupAutoCompleter(shell),
		InterruptPrompt:   "^C",
		HistorySearchFold: true,
	}
	rl, err := readline.NewEx(config)

	if err != nil {
		return err
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil { // io.EOF
			break
		}

		fmt.Println("")

		shell.ProcessLine(line)

		fmt.Println("")
	}
	return nil
}

func (shell *Shell) ProcessLine(line string) {
	tokens := strings.Split(line, " ")
	var command string
	var args = []string{}
	if len(tokens) == 0 {
		command = line
	} else {
		command = tokens[0]
		args = tokens[1:]
	}

	for _, cmd := range shell.Commands {
		if strings.ToLower(cmd.Name) == strings.ToLower(command) {
			cmd.Call(shell, args)
		}
	}
}

func (sh *Shell) ShowHelp() {
	for _, cmd := range sh.Commands {
		fmt.Printf("%s - %s\n", cmd.Name, cmd.Description)
	}
}
