package clay

import (
	"fmt"
)

type App struct {
	name        string
	description string
	usage       string
	version     string
	feature     Feature
}

type Action struct {
	handler       func()
	isRequiredArg bool
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
func (a *App) Short(flag string, handler func(), isRequiredArg bool) {
	if flag == "" {
		return
	}
	flag = fmt.Sprintf("-%s", flag)
	action := Action{
		handler:       handler,
		isRequiredArg: isRequiredArg,
	}
	a.feature[flag] = action
}

// This is for long flag. -> "--help", "--version"
func (a *App) Long(flag string, handler func(), isRequiredArg bool) {
	if flag == "" {
		return
	}
	flag = fmt.Sprintf("--%s", flag)
	action := Action{
		handler:       handler,
		isRequiredArg: isRequiredArg,
	}
	a.feature[flag] = action
}

// This is for parsing arguments
func (a *App) Parse() {
	// TODO: parsing arguments, and run the action handler
}
