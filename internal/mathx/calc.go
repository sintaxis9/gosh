package mathx

import (
	"fmt"
	"go/constant"
	"go/parser"
	"go/token"
	"go/types"
)

func Eval(expr string) (float64, error) {
	fs := token.NewFileSet()
	_, err := parser.ParseExpr(expr)
	if err != nil {
		return 0, fmt.Errorf("syntax error: %v", err)
	}

	tv, err := types.Eval(fs, nil, token.NoPos, expr)
	if err != nil {
		return 0, fmt.Errorf("evaluation error:  %v", err)
	}

	val := tv.Value
	if val == nil {
		return 0, fmt.Errorf("invalid expression")
	}

	f, _ := constant.Float64Val(val)
	return f, nil
}
