package command

import (
	"das-frama/dudes-bot/pkg/bot"
	"fmt"
)

// Response is a result of command process.
type Response struct {
	Text  string
	Photo string
}

var commandMap = map[string]func(bot.Update) Response{
	"start":     start,
	"stop":      stop,
	"ping":      ping,
	"schedule":  schedule,
	"call":      call,
	"overwatch": overwatch,
	"cat":       cat,
}

// Process handles the command and returns a response struct.
func Process(cmd string, update bot.Update) (Response, error) {
	fn, ok := commandMap[cmd]
	if !ok {
		return Response{}, fmt.Errorf("wrong command")
	}

	return fn(update), nil
}
