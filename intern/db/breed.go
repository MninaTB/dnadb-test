package db

import (
	"database/sql"

	"github.com/MninaTB/dnadb/intern/model"
)

const (
	selectBreed     = `SELECT id, name, mail, internet, country FROM breed`
	selectBreedByID = selectBreed + ` WHERE id = ?`
	deleteBreedByID = `DELETE FROM breed WHERE id = ?`
	insertBreed     = `INSERT INTO breed (id, name, mail, internet, country)
					   VALUES(?,?,?,?,?)`
)

// NewBreedRepo - TODO
func NewBreedRepo(db *sql.DB) *BreedRepo {
	return &BreedRepo{db: db}
}

// BreedRepo - TODO
type BreedRepo struct {
	db *sql.DB
}

// GetByID - TODO
func (b *BreedRepo) GetByID(id string) (*model.Breed, error) {
	row := b.db.QueryRow(selectBreedByID, id)
	breed := &model.Breed{}
	err := row.Scan(&breed.ID, &breed.Name, &breed.Mail, &breed.Internet, &breed.Country)
	if err != nil {
		return nil, err
	}
	return breed, nil
}

// List - TODO
func (b *BreedRepo) List() ([]*model.Breed, error) {
	rows, err := b.db.Query(selectBreed)
	if err != nil {
		return nil, err
	}
	var breeds []*model.Breed
	for rows.Next() {
		breed := &model.Breed{}
		err := rows.Scan(&breed.ID, &breed.Name, &breed.Mail, &breed.Internet, &breed.Country)
		if err != nil {
			return nil, err
		}
		breeds = append(breeds, breed)
	}
	return breeds, nil
}

// Create - TODO
func (b *BreedRepo) Create(breed *model.Breed) error {
	_, err := b.db.Exec(insertBreed, breed.ID, breed.Name, breed.Mail, breed.Internet, breed.Country)
	return err
}

// Delete - TODO
func (b *BreedRepo) Delete(id string) error {
	_, err := b.db.Exec(deleteBreedByID, id)
	return err
}
