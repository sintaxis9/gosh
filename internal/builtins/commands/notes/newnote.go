package notes

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func NewNote(args []string) {
	if len(args) < 2 {
		fmt.Println("usage: newnote <topic> \"content\"")
		return
	}

	topic := args[0]
	note := strings.Join(args[1:], " ")
	note = strings.Trim(note, "\"")

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("error getting home:", err)
		return
	}

	notesDir := filepath.Join(home, "Documents", "notes")
	os.MkdirAll(notesDir, 0755)

	filePath := filepath.Join(notesDir, topic+".txt")
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open file error", err)
		return
	}
	defer f.Close()

	info, _ := f.Stat()
	if info.Size() > 0 {
		f.WriteString("\n\n")
	}

	f.WriteString(fmt.Sprintf("-> \"%s\"", note))
	fmt.Printf("note append in %s\n", filePath)
}
