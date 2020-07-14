package command

import (
	"fmt"
)

// Result is a result of command process.
type Result struct {
	Text  string
	Photo string
}

var commandMap = map[string]func(params []string) (Result, error){
	"start":     start,
	"stop":      stop,
	"ping":      ping,
	"schedule":  schedule,
	"call":      call,
	"overwatch": overwatch,
	"cat":       cat,
}

// Process handles the command and returns a response struct.
func Process(cmd string, params []string) (Result, error) {
	fn, ok := commandMap[cmd]
	if !ok {
		return Result{}, fmt.Errorf("wrong command")
	}

	return fn(params)
}
