package notes

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func ShowNote(args []string) {
	if len(args) < 1 {
		fmt.Println("usage: shownote <topic> [last|first|index]")
		return
	}

	topic := args[0]
	home, _ := os.UserHomeDir()
	notesDir := filepath.Join(home, "Documents", "notes")
	filePath := filepath.Join(notesDir, topic+".txt")

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("error opening note %s: %v\n", topic, err)
		return
	}
	defer f.Close()

	notes := []string{}
	scanner := bufio.NewScanner(f)
	var current strings.Builder
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" && current.Len() > 0 {
			notes = append(notes, strings.TrimSpace(current.String()))
			current.Reset()
		} else {
			if current.Len() > 0 {
				current.WriteString("\n")
			}
			current.WriteString(line)
		}
	}
	if current.Len() > 0 {
		notes = append(notes, strings.TrimSpace(current.String()))
	}

	if len(args) == 1 {
		fmt.Printf("========== %s ==========\n", filepath.Base(filePath))
		for i, note := range notes {
			fmt.Printf("[%d]\n%s\n\n", i, note)
		}
		fmt.Println("===========================")
		return
	}

	option := args[1]
	switch option {
	case "last":
		fmt.Printf("[last]\n%s\n", notes[len(notes)-1])
	case "first":
		fmt.Printf("[first]\n%s\n", notes[0])
	default:
		index, err := strconv.Atoi(option)
		if err != nil {
			fmt.Println("invalid option, use: last, first, or index number")
			return
		}
		if index < 0 || index >= len(notes) {
			fmt.Printf("error: note index %d out of range (0-%d)\n", index, len(notes)-1)
			return
		}
		fmt.Printf("[%d]\n%s\n", index, notes[index])
	}
}
