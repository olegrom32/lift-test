package service

import (
	"errors"
	"log"
	"time"

	"testttt/internal/domain"
)

var ErrInvalidFloor = errors.New("invalid floor")

type Config struct {
	NumFloors int
}

type Brain struct {
	lift              *domain.Lift
	ticker            *time.Ticker
	currentFloor      int
	nextMoveDirection *bool // 0 - down, 1 - up
	floorsToVisit     int

	activeFloors []bool
}

func NewBrain(cfg Config, l *domain.Lift) *Brain {
	s := &Brain{
		lift:         l,
		ticker:       time.NewTicker(time.Second),
		activeFloors: make([]bool, cfg.NumFloors-1),
	}

	go func() {
		if err := s.tick(); err != nil {
			log.Print(err)
		}
	}()

	return s
}

func (s *Brain) tick() error {
	for {
		<-s.ticker.C

		switch {
		case s.lift.DoorsOpen:
			log.Print("Lift closing doors")

			if err := s.lift.CloseDoors(); err != nil {
				log.Printf("Failed to close the doors: %s", err.Error())
			}

			continue
		case s.nextMoveDirection == nil:
			continue
		case *s.nextMoveDirection == false:
			s.currentFloor--

			s.arrive()
		case *s.nextMoveDirection == true:
			s.currentFloor++

			s.arrive()
		}
	}
}

func (s *Brain) CallLift(n int) error {
	if n == s.currentFloor {
		//
	}

	if n < 0 || n >= len(s.activeFloors) {
		return ErrInvalidFloor
	}

	s.activeFloors[n] = true
	s.floorsToVisit++

	if s.nextMoveDirection == nil {
		if n < s.currentFloor {
			s.nextMoveDirection = &[]bool{false}[0]
		} else {
			s.nextMoveDirection = &[]bool{true}[0]
		}
	}

	return nil
}

func (s *Brain) arrive() {
	log.Printf("Lift arrived at floor: %d", s.currentFloor)

	if s.activeFloors[s.currentFloor] {
		if err := s.lift.OpenDoors(); err != nil {
			log.Printf("Doors won't open: %s", err.Error())
		}

		s.activeFloors[s.currentFloor] = false
		s.floorsToVisit--
	}

	if s.floorsToVisit == 0 {
		s.nextMoveDirection = nil
	}
}

func (s *Brain) setActiveFloor(n int, v bool) error {

	return nil
}
