package db

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"

	"github.com/MninaTB/dnadb/intern/model"
)

const (
	selectDog     = `SELECT id, name, birthdate, race, color, mother, father FROM dog`
	selectDogByID = selectDog + ` WHERE id = ?`
	deleteDogByID = `DELETE FROM dog WHERE id = ?`
	insertDog     = `INSERT INTO dog (id, name, birthdate, race, color, mother, father) 
					 VALUES(?,?,?,?,?,?,?)`
)

// NewDogRepo - TODO
func NewDogRepo(db *sql.DB) *DogRepo {
	return &DogRepo{db: db}
}

// DogRepo - TODO
type DogRepo struct {
	db *sql.DB
}

// GetByID - TODO
func (d *DogRepo) GetByID(id string) (*model.Dog, error) {
	row := d.db.QueryRow(selectDogByID, id)
	dog := &model.Dog{}
	var birthdate mysql.NullTime
	err := row.Scan(&dog.ID, &dog.Name, &birthdate, &dog.Race, &dog.Color, &dog.Mother, &dog.Father)
	if err != nil {
		return nil, err
	}
	if birthdate.Valid {
		dog.Brithdate = &birthdate.Time
	}
	return dog, nil
}

// List - TODO
func (d *DogRepo) List() ([]*model.Dog, error) {
	rows, err := d.db.Query(selectDog)
	if err != nil {
		return nil, err
	}

	var dogs []*model.Dog
	for rows.Next() {
		dog := &model.Dog{}
		var birthdate mysql.NullTime
		err := rows.Scan(&dog.ID, &dog.Name, &birthdate, &dog.Race, &dog.Color, &dog.Mother, &dog.Father)
		if err != nil {
			return nil, err
		}
		if birthdate.Valid {
			dog.Brithdate = &birthdate.Time
		}
		dogs = append(dogs, dog)
	}
	return dogs, nil
}

// Create - TODO
func (d *DogRepo) Create(dog *model.Dog) error {
	_, err := d.db.Exec(insertDog, dog.ID, dog.Name, dog.Brithdate, dog.Race, dog.Color, dog.Mother, dog.Father)
	return err
}

// Delete - TODO
func (d *DogRepo) Delete(id string) error {
	_, err := d.db.Exec(deleteDogByID, id)
	return err
}
