package ascii

func StartIndex(c rune) int {
	return int((c - 32) + 1 + 8*(c-32))
}

/*
res := Args{}
	res.banner = "standard"
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if strings.HasPrefix(arg, "--output=") {
			if res.output != "" {
				return res, errOuputAlreadyDefiend
			}
			if res.input != "" {
				return res, errUsage
			}
			if len(arg) <= 9 {
				return res, errOutputFileNotDefined
			}
			filename := arg[9:]
			fmt.Println(filename)
			res.output = filename
		} else if res.input == "" {
			res.input = arg
		} else {
			res.banner = arg
		}
	}
	if res.input == "" {
		return res, errInputNotDefiend
	}
	return res, nil
*/
