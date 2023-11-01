package domain

import (
	"errors"
	"math/rand"
)

var ErrDoorsBlocked = errors.New("doors blocked")

type Lift struct {
	DoorsOpen bool
}

func NewLift() *Lift {
	return &Lift{
		DoorsOpen: false,
	}
}

func (l *Lift) OpenDoors() error {
	if rand.Int()%2 == 0 {
		return ErrDoorsBlocked
	}

	l.DoorsOpen = true

	return nil
}

func (l *Lift) CloseDoors() error {
	if rand.Int()%2 == 0 {
		return ErrDoorsBlocked
	}

	l.DoorsOpen = false

	return nil
}
