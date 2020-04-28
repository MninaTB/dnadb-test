package db

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"

	"github.com/MninaTB/dnadb/intern/model"
)

const (
	selectAccount     = `SELECT id, mail, password, nickname, firstname, lastname, country, street, streetnumber, zipcode, city, breeder, created, updated FROM account`
	selectAccountByID = selectAccount + ` WHERE id = ?`
	deleteAccountByID = `DELETE FROM account WHERE id = ?`
	insertAccount     = `INSERT INTO account (id, mail, password, nickname, firstname, lastname, country, street, streetnumber, zipcode, city, breeder, created, updated)
						 Values(?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
)

// NewAccountRepo - TODO
func NewAccountRepo(db *sql.DB) *AccountRepo {
	return &AccountRepo{db: db}
}

// AccountRepo - TODO
type AccountRepo struct {
	db *sql.DB
}

// GetByID - TODO
func (a *AccountRepo) GetByID(id string) (*model.Account, error) {
	row := a.db.QueryRow(selectAccountByID, id)
	account := &model.Account{}
	var created mysql.NullTime
	var updated mysql.NullTime
	err := row.Scan(&account.ID, &account.Mail, &account.Password, &account.Nickname, &account.Firstname, &account.Lastname, &account.Country,
		&account.Street, &account.Streetnumber, &account.Zipcode, &account.City, &account.Breeder, &created, &updated)
	if err != nil {
		return nil, err
	}
	if created.Valid {
		account.Created = &created.Time
	}
	if updated.Valid {
		account.Updated = &updated.Time
	}
	return account, nil
}

// List - TODO
func (a *AccountRepo) List() ([]*model.Account, error) {
	rows, err := a.db.Query(selectAccount)
	if err != nil {
		return nil, err
	}

	var accounts []*model.Account
	for rows.Next() {
		account := &model.Account{}
		var created mysql.NullTime
		var updated mysql.NullTime
		err := rows.Scan(&account.ID, &account.Mail, &account.Password, &account.Nickname, &account.Firstname, &account.Lastname, &account.Country,
			&account.Street, &account.Streetnumber, &account.Zipcode, &account.City, &account.Breeder, &created, &updated)
		if err != nil {
			return nil, err
		}
		if created.Valid {
			account.Created = &created.Time
		}
		if updated.Valid {
			account.Updated = &updated.Time
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

//Create - TODO
func (a *AccountRepo) Create(account *model.Account) error {
	_, err := a.db.Query(insertAccount, account.ID, account.Mail, account.Password, account.Nickname, account.Firstname, account.Lastname, account.Country,
		account.Street, account.Streetnumber, account.Zipcode, account.City, account.Breeder, account.Created, account.Updated)
	return err
}

//Delete - TODO
func (a *AccountRepo) Delete(account *model.Account) error {
	_, err := a.db.Query(deleteAccountByID, account.ID, account.Mail, account.Password, account.Nickname, account.Firstname, account.Lastname, account.Country,
		account.Street, account.Streetnumber, account.Zipcode, account.City, account.Breeder, account.Created, account.Updated)
	return err
}
