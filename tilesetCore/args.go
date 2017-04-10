package main

import (
	"errors"
	"os"
	"strings"
)

const (
	INPUT      = "-i"
	OUTPUT     = "-o"
	CORE       = "-c"
	MODIFY_HUI = "-m"
)

type Args map[string]string

var (
	defaultConf = Args{
		CORE:       "cpu",
		MODIFY_HUI: "false",
	}
	requireArg = []string{INPUT, OUTPUT}
)

func parseArgs() (*Args, error) {
	var (
		err error
	)
	arg := make(Args)
	for _, v := range os.Args {
		d := strings.Split(v, "=")
		if len(d) == 2 {
			arg[d[0]] = d[1]
		}
	}
	for key, v := range defaultConf {
		if _, find := arg[key]; !find {
			arg[key] = v
		}
	}

	for _, v := range requireArg {
		if _, find := arg[v]; !find {
			err = errors.New("required data not completed")
			return nil, err
		}
	}
	return &arg, nil
}

func (a *Args) Get(key string) string {
	if v, find := (*a)[key]; find {
		return v
	}
	return ""
}
