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
	fmt.Println("create account repo")
	ar := database.NewAccountRepo(db)
	myaccount := &model.Account{
		ID:        "1",
		Mail:      "blub@bla.de",
		Password:  "2",
		Nickname:  "Silvia",
		Firstname: "Patrick",
		Lastname:  "Pups",
		Country:   model.AccountCountryGermany,
	}
	fmt.Println("ok")
	testaccount, err := ar.GetByID(string(myaccount.ID))
	if err == sql.ErrNoRows {
		fmt.Println("erzeugt Account")
		err = ar.Create(myaccount)
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("Account existiert", testaccount)
		return
	}
	fmt.Println("fetch account list")
	accounts, err := ar.List()
	if err != nil {
		panic(err)
	}
	for _, account := range accounts {
		fmt.Println(account)
	}
	fmt.Println("done")
}
