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

type DogColor int

const (
	DogColorUnknown DogColor = iota
	DogColorRed
	DogColorBrown
)

type UUID string

type Dog struct {
	ID        UUID
	Name      string
	Brithdate *time.Time
	Race      DogRace
	Color     DogColor
	Mother    UUID
	Father    UUID
}
