package clay

import (
	"fmt"
	"os"
)

type App struct {
	name        string
	description string
	usage       string
	version     string
	feature     Feature
}

type Action struct {
	handler interface{}
}

type Feature map[string]Action

func New(name, description, usage, version string) *App {
	feature := make(Feature)
	return &App{
		name:        name,
		description: description,
		usage:       usage,
		version:     version,
		feature:     feature,
	}
}

// This is for short flag. -> "-h", "-v"
func (a *App) Short(flag string, handler interface{}) {
	if flag == "" {
		return
	}
	flag = fmt.Sprintf("-%s", flag)
	action := Action{
		handler: handler,
	}
	a.feature[flag] = action
}

// This is for long flag. -> "--help", "--version"
func (a *App) Long(flag string, handler interface{}) {
	if flag == "" {
		return
	}
	flag = fmt.Sprintf("--%s", flag)
	action := Action{
		handler: handler,
	}
	a.feature[flag] = action
}

// This is for parsing arguments
func (a *App) Parse() {
	for i := 1; i < len(os.Args[1:]); i++ {
		flag := os.Args[i]
		feature, ok := a.feature[flag]
		if !ok {
			continue
		}

		switch h := feature.handler.(type) {
		case func():
			h()
		case func(string):
			h(os.Args[i+1])
		default:
			panic("Invalid Action handler")
		}
	}
}
