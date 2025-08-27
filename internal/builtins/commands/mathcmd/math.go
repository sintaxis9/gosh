package mathcmd

import (
	"fmt"
	"strconv"

	"github.com/sintaxis9/gosh/internal/mathx"
)

func Math(args []string) {
	if len(args) == 0 {
		fmt.Println("usage: math <eval|sin|cos|tan> <expression|number>")
		return
	}

	switch args[0] {
	case "eval":
		if len(args) < 2 {
			fmt.Println("usage: math eval <expression>")
			return
		}
		result, err := mathx.Eval(args[1])
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		fmt.Println(result)

	case "sin", "cos", "tan":
		if len(args) < 2 {
			fmt.Printf("usage: math %s <number>\n", args[0])
			return
		}
		val, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
			fmt.Println("invalid number:", args[1])
			return
		}

		switch args[0] {
		case "sin":
			fmt.Println(mathx.Sin(val))
		case "cos":
			fmt.Println(mathx.Cos(val))
		case "tan":
			fmt.Println(mathx.Tan(val))
		}

	default:
		fmt.Println("unknown math subcommand:", args[0])
	}
}
