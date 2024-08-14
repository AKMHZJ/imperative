package ascii

import (
	"errors"
	"strings"
)

type Args struct {
	Banner string
	Output string
	Input  string
}

var (
	errOutputFileNotDefined = errors.New("output file is not defined")
	errOuputAlreadyDefiend  = errors.New("output argument is already defined")
	errUsage                = errors.New("usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
	errInputNotDefiend      = errors.New("input string not defiend")
	errIncorrectedFilename  = errors.New("usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
)

func HandleArgs(args []string) (Args, error) {
	res := Args{}
	isInputDefined := false
	res.Banner = "standard"
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if strings.HasPrefix(arg, "--output=") {
			if res.Output != "" {
				return res, errOuputAlreadyDefiend
			}
			if res.Input != "" {
				return res, errUsage
			}
			if len(arg) <= 9 {
				return res, errOutputFileNotDefined
			}
			filename := arg[9:]
			if !strings.HasSuffix(filename, ".txt") {
				return res, errIncorrectedFilename
			}
			res.Output = filename
		} else if res.Input == "" && !isInputDefined {
			res.Input = arg
			isInputDefined = true
		} else {
			res.Banner = arg
		}
	}
	if res.Input == "" {
		return res, errInputNotDefiend
	}
	return res, nil
}
