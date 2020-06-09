package command

import "fmt"

type CommandResponse struct {
	Reply string
}

var commandMap = map[string]func(string) CommandResponse{
	"start":    start,
	"ping":     ping,
	"schedule": schedule,
	"call":     call,
}

func Process(cmd string, params string) (string, error) {
	if fn, ok := commandMap[cmd]; ok {
		return fn(params).Reply, nil
	}

	return "", fmt.Errorf("wrong command")
	// switch cmd {
	// case "start":
	// 	reply = "А я уже запущен."
	// case "ping":
	// 	reply = "pong"
	// case "joke":
	// 	reply = "Если ваш крем от морщин действует, то почему у вас до сих пор есть отпечатки пальцев?"
	// case "judge":
	// 	reply = "Хороший выбор! Однако я ещё не научился быть как Дредд, поэтому пускай будет прав тот, кто первым меня об этом попросил."
	// case "advice":
	// 	reply = "У меня есть только один совет: не следует множить сущее без необходимости."
	// case "vacation":
	// 	reply = "Сегодня Саша трудится в поте лица"
	// }
}
