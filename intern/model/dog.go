package model

import (
	"time"
)

type DogRace int

const (
	DogRaceUnknown DogRace = iota
	DogRaceGorilla
	DogRaceSchnuffi
)

type UUID string

type Dog struct {
	ID        UUID
	Name      string
	Brithdate *time.Time
	Race      DogRace
	Mother    UUID
	Father    UUID
}
