package filesystem

import (
	"fmt"
	"os"
	"time"

	"github.com/sintaxis9/gosh/internal/utils"
)

func Ls(args []string) {
	path := "."
	showDetails := false
	showHidden := false

	var filteredArgs []string
	for _, arg := range args {
		switch arg {
		case "-l":
			showDetails = true
		case "-a":
			showHidden = true
		default:
			filteredArgs = append(filteredArgs, arg)
		}
	}

	if len(filteredArgs) > 0 {
		path = filteredArgs[0]
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("error reading directory %s: %v\n", path, err)
		return
	}

	for _, entry := range entries {
		name := entry.Name()

		if !showHidden && utils.IsHidden(entry) {
			continue
		}

		if showDetails {
			info, err := entry.Info()
			if err != nil {
				fmt.Printf("error getting info for %s\n", name)
				continue
			}

			perm := info.Mode().Perm()
			size := info.Size()
			modTime := info.ModTime().Format(time.RFC822)

			if entry.IsDir() {
				fmt.Printf("%s [DIR]  %-20s %10d bytes  %s\n", perm, name, size, modTime)
			} else {
				fmt.Printf("%s [FILE] %-20s %10d bytes  %s\n", perm, name, size, modTime)
			}

		} else {
			if entry.IsDir() {
				fmt.Printf("\033[34m%s\033[0m\n", name)
			} else {
				fmt.Println(name)
			}
		}
	}
}
