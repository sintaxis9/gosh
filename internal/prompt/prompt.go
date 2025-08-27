package prompt

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func Get() string {
	// USER
	usr, err := user.Current()
	username := "user"
	if err == nil {
		parts := strings.Split(usr.Username, `\`)
		username = parts[len(parts)-1]
	}

	// HOSTNAME
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "host"
	}

	// FOLDER
	cwd, err := os.Getwd()
	currentFolder := "?"
	if err == nil {
		currentFolder = filepath.Base(cwd)
	}

	return fmt.Sprintf("[%s@%s:%s]> ", hostname, username, currentFolder)
}
