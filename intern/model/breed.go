package model

type BreedCountry int

const (
	BreedCountryGermany BreedCountry = iota
	BreedCountryAustria
)

type Breed struct {
	ID       UUID
	Name     string
	Mail     string
	Internet string
	Country  BreedCountry
}
