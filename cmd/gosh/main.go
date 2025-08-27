package main

import (
	"github.com/sintaxis9/gosh/internal/builtins"
	"github.com/sintaxis9/gosh/internal/prompt"
	"github.com/sintaxis9/gosh/internal/shell"
)

// usage of shell.go
func main() {
	s := shell.NewShell(prompt.Get, builtins.Commands)
	s.Run()
}
