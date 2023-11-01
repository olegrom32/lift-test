package main

import (
	"log"
	"time"

	"testttt/internal/domain"
	"testttt/internal/service"
)

func main() {
	lift := domain.NewLift()

	brain := service.NewBrain(service.Config{NumFloors: 10}, lift)

	if err := brain.CallLift(3); err != nil {
		log.Print(err)
	}

	time.Sleep(time.Minute)
}
