package main

import (
	"database/sql"
	"fmt"

	"github.com/MninaTB/dnadb/intern/model"

	_ "github.com/go-sql-driver/mysql"

	database "github.com/MninaTB/dnadb/intern/db"
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
		ID:    "3",
		Name:  "Klaus",
		Color: model.DogColorBrown,
		Race:  model.DogRaceGorilla,
	}
	//err = dr.Delete(string(mydog.ID))
	if err != nil {
		panic(err)
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
		fmt.Println("Hund existiert", testdog)
		return
	}
	fmt.Println("fetch dog list")
	dogs, err := dr.List()
	if err != nil {
		panic(err)
	}
	for _, dog := range dogs {
		fmt.Println(dog)
	}
	fmt.Println("done")
}
