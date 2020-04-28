package model

import "time"

type AccountCountry int

const (
	AccountCountryNetherlands AccountCountry = iota
	AccountCountryGermany
	AccountCountryAustria
)

type Account struct {
	ID           UUID
	Mail         string
	Password     string
	Nickname     string
	Firstname    string
	Lastname     string
	Country      AccountCountry
	Street       string
	Streetnumber int
	Zipcode      int
	City         string
	Breeder      bool
	Created      *time.Time
	Updated      *time.Time
}
