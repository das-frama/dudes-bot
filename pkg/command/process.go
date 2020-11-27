package command

import (
	"das-frama/dudes-bot/pkg/bot"
	"das-frama/dudes-bot/pkg/config"
	"das-frama/dudes-bot/pkg/sqlite"
	"fmt"
)

// Config struct
type commandConfig struct {
	Message   *bot.Message
	Queryer   sqlite.Queryer
	AppConfig config.Config
}

// Result is a result of command process.
type Result struct {
	Text      string
	PhotoURL  string
	PhotoData []byte
}

var commandMap = map[string]func(commandConfig) (Result, error){
	"start":     start,
	"stop":      stop,
	"help":      help,
	"ping":      ping,
	"schedule":  schedule,
	"call":      call,
	"overwatch": overwatch,
	"cat":       cat,
	"dog":       dog,
	"panda":     panda,
	"meme":      meme,
	"poetry":    poetry,
}

// Process handles the command and returns a response struct.
func Process(cmd string, m *bot.Message, q sqlite.Queryer, cfg config.Config) (Result, error) {
	fn, ok := commandMap[cmd]
	if !ok {
		return Result{}, fmt.Errorf("wrong command")
	}

	return fn(commandConfig{
		Message:   m,
		Queryer:   q,
		AppConfig: cfg,
	})
}
