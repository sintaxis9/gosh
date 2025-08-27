package builtins

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/sintaxis9/gosh/internal/builtins/commands/fs"
	"github.com/sintaxis9/gosh/internal/builtins/commands/mathcmd"
	"github.com/sintaxis9/gosh/internal/builtins/commands/notes"
	"github.com/sintaxis9/gosh/internal/builtins/commands/system"
)

func Exit(args []string) {
	println("xau")
	os.Exit(0)
}

func Clear(args []string) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

var Commands = map[string]func(args []string){
	// DEFAULT
	"exit":  Exit,
	"clear": Clear,

	// FILESYSTEM
	"cd":    fs.Cd,
	"ls":    fs.Ls,
	"mkdir": fs.Mkdir,
	"rmdir": fs.Rmdir,
	"touch": fs.Touch,
	"rm":    fs.Rm,

	// SYSTEM
	"sysinfo": system.Sysinfo,

	// NOTES
	"newnote":  notes.NewNote,
	"shownote": notes.ShowNote,
	"delnote":  notes.DelNote,

	// MATH
	"math": mathcmd.Math,
}
