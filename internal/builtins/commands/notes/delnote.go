package notes

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func DelNote(args []string) {
	if len(args) != 2 {
		fmt.Println("usage: delnote <topic> [index]")
		return
	}

	topic := args[0]
	indexStr := args[1]
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		fmt.Println("error: index is a num")
		return
	}

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("error getting home dir:", err)
		return
	}

	notesDir := filepath.Join(home, "Documents", "notes")
	filePath := filepath.Join(notesDir, topic+".txt")

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("error opening note %s: %v\n", topic, err)
		return
	}
	defer f.Close()

	var notes []string
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

	if index < 0 || index >= len(notes) {
		fmt.Printf("error: note index %d out of range (0-%d)\n", index, len(notes)-1)
		return
	}

	removed := notes[index]
	notes = append(notes[:index], notes[index+1:]...)

	newContent := strings.Join(notes, "\n\n")
	err = os.WriteFile(filePath, []byte(newContent), 0644)
	if err != nil {
		fmt.Println("error rewriting file:", err)
		return
	}

	fmt.Printf("note [%d] exterminated from %s\n", index, filePath)
	fmt.Printf("deleted content:\n-> %s\n", removed)
}
