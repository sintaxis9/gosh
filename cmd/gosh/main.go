package main

import (
	"github.com/sintaxis9/gosh/internal/builtins"
	"github.com/sintaxis9/gosh/internal/prompt"
	"github.com/sintaxis9/gosh/internal/shell"
)

func main() {
	s := shell.NewShell(prompt.Get, builtins.Commands)
	s.Run()
}
