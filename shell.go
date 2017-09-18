package gshell

import (
	"fmt"
	"github.com/chzyer/readline"
	"sort"
	"strings"
)

type Shell struct {
	Commands     []*Command
	Prompt       string
	VimMode      bool
	HistoryLimit int
}

type ByName []*Command

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func New() *Shell {
	sh := &Shell{
		Prompt:       "> ",
		Commands:     []*Command{},
		HistoryLimit: 500,
		VimMode:      false,
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
		Prompt:          shell.Prompt,
		AutoComplete:    setupAutoCompleter(shell),
		InterruptPrompt: "^C",
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
			break
		}
	}
}

func (sh *Shell) ShowHelp() {
	sort.Sort(ByName(sh.Commands))
	for _, cmd := range sh.Commands {
		fmt.Printf("%s - %s\n", cmd.Name, cmd.Description)
	}
}
