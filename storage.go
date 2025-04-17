package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

//This is called REPOSITORY LAYER INTERFACE
//interface of the storage -- Defining a contract for any datbase
//This is an interface that defines the required methods for any kind of storage backend to work with Account.

//-----------WHY USE THIS? --------------------
/*Abstraction: Your business logic doesn’t care if you're using PostgreSQL, MongoDB, or even an in-memory store.

Decoupling: Keeps your core logic independent of infrastructure details.
*/

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)
	GetAccountByID(int) (*Account, error)
}

// Concrete Storage Implimentation
// Holds the DB connection Objects
type PostgresStore struct {
	db *sql.DB
}

// Constructor
// basically again a constructor for the PostgresStore were we are creating the db with sql open and using the postgres as the driver for that.
// we are opening a query to the database via that and creating a connecting string
// later we ping the data to check the connection is still alive or not
// the return the object ( the Postgres connection )
func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=goBank sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

// Initalizer
// An init fucntion to run before anything in the file
// This Ensures that the tables exist before using the DB.
func (s *PostgresStore) Init() error {
	return s.CreateAccountTable()

}

// these are the function of the interface of the structure Postgres and here pass the query to the database
// create the datbase table schema
func (s *PostgresStore) CreateAccountTable() error {
	query := ` create table if not exists account(
		id serial primary key,
		first_name varchar(50),
		last_name varchar(50),
		number serial,
		balance serial,
		created_at timestamp
	
	)`

	_, err := s.db.Exec(query)
	return err
}

// To create a new account
// Insert opearation
func (s *PostgresStore) CreateAccount(acc *Account) error {
	query := `insert into account
		(first_name,last_name,number,balance,created_at)
	    values ($1,$2,$3,$4,$5)
	`

	resp, err := s.db.Exec(query, acc.FirstName, acc.LastName, acc.Number, acc.Balance, acc.CreatedAt)
	if err != nil {
		return err
	}
	fmt.Println(resp)

	return nil
}
func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}

func (s *PostgresStore) DeleteAccount(int) error {
	return nil
}

func (s *PostgresStore) GetAccountByID(int) (*Account, error) {
	return nil, nil
}

func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	rows, err := s.db.Query(`select * from account`)
	if err != nil {
		return nil, err
	}

	accounts := []*Account{}

	for rows.Next() {
		account, err := scanIntoAccount(rows)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil

}

func scanIntoAccount(rows *sql.Rows) (*Account, error) {
	account := new(Account)

	err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Number,
		&account.Balance,
		&account.CreatedAt)

	return account, err
}
