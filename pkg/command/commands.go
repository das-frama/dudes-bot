package command

import (
	"das-frama/dudes-bot/pkg/bot"
	"fmt"
)

type CommandResponse struct {
	Reply string
}

var commandMap = map[string]func(bot.Update) CommandResponse{
	"start":     start,
	"ping":      ping,
	"schedule":  schedule,
	"call":      call,
	"overwatch": overwatch,
}

func Process(cmd string, update bot.Update) (string, error) {
	if fn, ok := commandMap[cmd]; ok {
		return fn(update).Reply, nil
	}

	return "", fmt.Errorf("wrong command")
}
