package command

import (
	"log"
	"time"
)

const firstWorkDay = "05.06.2020"

func start() CommandResponse {
	return CommandResponse{
		Reply: "А я уже запущен",
	}
}

func ping() CommandResponse {
	return CommandResponse{
		Reply: "pong",
	}
}

func schedule() CommandResponse {
	workTime, err := time.Parse("02.01.2006", firstWorkDay)
	if err != nil {
		log.Fatalln(err)
	}
	currentTime := time.Now()
	hours := currentTime.Sub(workTime).Hours()
	days := int(hours / 24)
	var reply string
	if days%4 < 2 {
		reply = "Сегодня Саша трудится в поте лица!"
	} else {
		reply = "Сегодня Саша отдыхает! 😊😊😊"
	}

	return CommandResponse{
		Reply: reply,
	}
}
