package command

import (
	"fmt"
	"log"
	"strings"
	"time"
)

const firstWorkDay = "05.06.2020"

func start(msg string) CommandResponse {
	return CommandResponse{
		Reply: "А я уже запущен",
	}
}

func ping(msg string) CommandResponse {
	return CommandResponse{
		Reply: "pong",
	}
}

func schedule(msg string) CommandResponse {
	workTime, err := time.Parse("02.01.2006", firstWorkDay)
	if err != nil {
		log.Fatalln(err)
	}

	var verb string
	if msg == "" {
		verb = "Сегодня"
	} else {
		verb = msg
	}

	currentTime := time.Now()
	switch msg {
	case "завтра":
		currentTime = currentTime.Add(time.Hour * 24)
	case "послезавтра":
		currentTime = currentTime.Add(time.Hour * 48)
	}

	hours := currentTime.Sub(workTime).Hours()
	days := int(hours / 24)

	var reply string
	if days%4 < 2 {
		reply = "%s Саша трудится в поте лица!"
	} else {
		reply = "%s Саша отдыхает! 😊😊😊"
	}

	return CommandResponse{
		Reply: fmt.Sprintf(reply, strings.Title(verb)),
	}
}
