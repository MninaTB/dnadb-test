package db

import (
	"database/sql"

	"github.com/MninaTB/dnadb/intern/model"
)

const (
	selectDog     = `SELECT id, name, birthdate, race, color, mother, father FROM dog`
	selectDogByID = selectDog + ` WHERE id = ?`
)

func NewDogRepo(db *sql.DB) *DogRepo {
	return &DogRepo{db: db}
}

type DogRepo struct {
	db *sql.DB
}

func (d *DogRepo) GetByID(id string) (*model.Dog, error) {
	rows, err := d.db.Query(selectDogByID, id)
	if err != nil {
		return nil, err
	}
	dog := &model.Dog{}
	return dog, rows.Scan(&dog.ID, &dog.Name, &dog.Race, &dog.Color, &dog.Mother, &dog.Father)
}

func (d *DogRepo) List() ([]*model.Dog, error) {
	rows, err := d.db.Query(selectDog)
	if err != nil {
		return nil, err
	}
	var dogs []*model.Dog
	for rows.Next() {
		dog := &model.Dog{}
		err := rows.Scan(&dog.ID, &dog.Name, &dog.Race, &dog.Color, &dog.Mother, &dog.Father)
		if err != nil {
			return nil, err
		}
		dogs = append(dogs, dog)
	}
	return dogs, nil
}

func (d *DogRepo) Create(dog *model.Dog) error {
	return nil
}

func (d *DogRepo) Delete(id string) error {
	return nil
}
