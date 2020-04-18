package model

import (
  "time"
)

type DogRace int

const (
  unknown DogRace = iota
  gorilla
  schnuffi
)

type DogColor int

const (
  unknown DogColor = iota
  red
  brown
)

type UUID string

type Dog struct {
  ID        UUID
  Name      string
  Brithdate *time.Time
  Race      DogRace
  Color     DogColor
  mother    UUID
  father    UUID
}
