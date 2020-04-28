package main

import (
	"database/sql"
	"fmt"

	database "github.com/MninaTB/dnadb/intern/db"
	"github.com/MninaTB/dnadb/intern/model"
)

func main() {
	fmt.Println("open db connection")
	db, err := sql.Open("mysql", "dnadb:1234@tcp(127.0.0.1:3306)/dnadb")
	if err != nil {
		panic(err)
	}

	fmt.Println("create dog repo")
	dr := database.NewDogRepo(db)
	mydog := &model.Dog{
		ID:   "8",
		Name: "Bene",
		Race: model.DogRaceUnknown,
	}
	testdog, err := dr.GetByID(string(mydog.ID))
	if err == sql.ErrNoRows {
		fmt.Println("erzeuge Hund")
		err = dr.Create(mydog)
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("Hund existiert bereits", testdog)
		return
	}

	fmt.Println("create breed repo")
	br := database.NewBreedRepo(db)
	mybreed := &model.Breed{
		ID:       "10",
		Name:     "Dream",
		Mail:     "dog@dream.breed",
		Internet: "www.dream.breed",
		Country:  model.BreedCountryAustria,
	}
	testbreed, err := br.GetByID(string(mybreed.ID))
	if err == sql.ErrNoRows {
		fmt.Println("erzeuge Zucht")
		err = br.Create(mybreed)
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("Zucht existiert bereits", testbreed)
		return
	}
	fmt.Println("repos created")

	fmt.Println("Liste Hunde")
	dogs, err := dr.List()
	if err != nil {
		panic(err)
	}
	for _, dog := range dogs {
		fmt.Println(dog)
	}

	fmt.Println("Liste Züchter")
	breeds, err := br.List()
	if err != nil {
		panic(err)
	}
	for _, breed := range breeds {
		fmt.Println(breed)
	}
	fmt.Println("done")
}
