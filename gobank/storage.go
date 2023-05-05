package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(int, *Account) error
	GetAccounts() ([]*Account, error)
	GetAccountByID(int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStore{
		db: db,
	}, nil
}

func (p *PostgresStore) CreateAccount(account *Account) error {
	_, err := p.db.Query(`insert into account 
	(firstName, secondName, number, balance, createdAt)
	values($1, $2, $3, $4, $5)`,
		account.FirstName,
		account.SecondName,
		account.Number,
		account.Balance,
		account.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresStore) DeleteAccount(id int) error {
	_, err := p.db.Query(`delete from account where id = $1`, id)
	return err
}

func (p *PostgresStore) UpdateAccount(id int, account *Account) error {
	return nil
}

func (p *PostgresStore) GetAccounts() ([]*Account, error) {
	rows, err := p.db.Query(`select * from account`)
	if err != nil {
		return nil, err
	}
	accounts := []*Account{}
	for rows.Next() {
		account, err := scanIntoAccounts(rows)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (p *PostgresStore) GetAccountByID(id int) (*Account, error) {
	rows, err := p.db.Query(`select * from account where id = $1`, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanIntoAccounts(rows)
	}
	return nil, fmt.Errorf("account with id %d not found", id)
}

func (s *PostgresStore) init() error {
	// err := s.DropAccountTable()
	// if err != nil {
	// 	return err
	// }
	err := s.createAccountTable()
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) createAccountTable() error {
	query := `create table if not exists account(
		id serial primary key,
		firstName varchar(64),
		secondName varchar(64),
		number integer,
		balance integer,
		createdAt timestamp
	);`
	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) DropAccountTable() error {
	_, err := s.db.Exec("drop table account;")
	return err
}

func scanIntoAccounts(rows *sql.Rows) (*Account, error) {
	account := new(Account)
	if err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.SecondName,
		&account.Number,
		&account.Balance,
		&account.CreatedAt,
	); err != nil {
		return nil, err
	}
	return account, nil
}
