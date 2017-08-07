## A simple shell style interface for console utilities

GShell is a simple library that makes it easy to rapidly build shell style interfaces. Complete with a number of helpful features.

* Command history
* Command history search
* Vim mode
* Easily handle arguments
* Tab commmand completion
* Automatically generated help function

## Installation

```
go get github.com/5sigma/gshell
```

## Usage

```go

import (
  "fmt"
  "github.com/5sigma/gshell"
)


// Create a new shell
shell := gshell.New()

// Add a command
shell.AddCommand(&gshell.Command{
  Name: "mycmd",
  Description: "My custom command description that appears in help",
  Call: func(shell *gshell, args []string) {
    fmt.Println("Hello from mycmd")
  });
})

shell.Start()
```
