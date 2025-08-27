package shell

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
			fmt.Println("Error leyendo entrada:", err)
			return
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		tokens := strings.Fields(line)
		cmd := tokens[0]
		args := tokens[1:]

		if cmdFunc, ok := s.Commands[cmd]; ok {
			cmdFunc(args)
		} else {
			fmt.Println("wth that command:", cmd)
		}
	}
}
