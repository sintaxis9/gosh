package shell

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sintaxis9/gosh/internal/parser"
	"github.com/sintaxis9/gosh/internal/plugins"
)

type Shell struct {
	Commands map[string]func(args []string)
	Prompt   func() string
}

func NewShell(promptFunc func() string, commands map[string]func(args []string)) *Shell {
	return &Shell{
		Commands: commands,
		Prompt:   promptFunc,
	}
}

func (s *Shell) Run() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(s.Prompt())

		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("reading input error:", err)
			return
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		cmd, args, err := parser.ParseCommand(line)
		if err != nil {
			fmt.Println("error parsing command:", err)
			continue
		}

		if cmd == "" {
			continue
		}

		if cmdFunc, ok := s.Commands[cmd]; ok {
			cmdFunc(args)
		} else {
			if plugin, ok := plugins.Get(cmd); ok {
				plugin.Run(args)
			} else {
				fmt.Println("wth that command:", cmd)
			}
		}

	}
}
