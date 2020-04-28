package main

import (
	"database/sql"
	"fmt"

	database "github.com/MninaTB/dnadb/intern/db"
	"github.com/MninaTB/dnadb/intern/model"
)

func main() {
	fmt.Println("open db")
	db, err := sql.Open("mysql", "dnadb:1234@tcp(127.0.0.1:3306)/dnadb")
	if err != nil {
		panic(err)
	}
	fmt.Println("create breed repo")
	br := database.NewBreedRepo(db)
	mybreed := &model.Breed{
		ID:       "2",
		Name:     "Miyu",
		Mail:     "miyu@miyu.miyu",
		Internet: "www.miyu.miyu",
		Country:  model.BreedCountryGermany,
	}
	fmt.Println("ok")
	testbreed, err := br.GetByID(string(mybreed.ID))
	if err == sql.ErrNoRows {
		fmt.Println("erzeugt Züchter")
		err = br.Create(mybreed)
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("Züchter exisitert", testbreed)
		return
	}
	fmt.Println("fetch breeder list")
	breeds, err := br.List()
	if err != nil {
		panic(err)
	}
	for _, breed := range breeds {
		fmt.Println(breed)
	}
	fmt.Println("done")
}
