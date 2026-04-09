package main

import (
	"database/sql"
	"fmt"
	// "log"

	_ "github.com/lib/pq" // To register the driver.
)

type Storage interface{
	CreateAccount(*Account) error
	// GetAccountByID(int) (*Account, error)
	// DeleteAccount(int) error
	// UpdateAccount(*Account) error	
}

type PostgresStorage struct{
	db *sql.DB
}

func newPostgresStorage()(*PostgresStorage, error){
	fmt.Println("Connecting to the database...")
	connStr := "user=postgres password=paritosh host=localhost port=5050 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("Error opening database: %v", err)
        return nil, err
    }

	var exists bool
	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = %s)", "$1")
	err = db.QueryRow(query, "bank").Scan(&exists)
	if err != nil {
    	fmt.Printf("Error checking database existence: %v", err)
		return nil, err
	}
	if exists {
		fmt.Println("Database already exists.")
	}else{
		db_name := "bank"
		query := fmt.Sprintf("CREATE DATABASE %s", db_name)
		log.
		_, err = db.Exec(query)
		if err != nil {
			fmt.Printf("Error creating database: %v", err)
			return nil, err
		}
	}

	if err := db.Ping(); err != nil {
		fmt.Printf("Error pinging database: %v", err)
		return nil, err
	}
	return &PostgresStorage{db: db}, nil
} 

func (ps *PostgresStorage) Init() error {
	return ps.CreateAccountTable()
}

func (ps *PostgresStorage) CreateAccountTable() error {
	fmt.Println("Creating accounts table if not exists...")
	query := `create table if not exists accounts (
		id serial primary key,
		first_name varchar(50),
		last_name varchar(50),
		number serial,
		balance float,
		created_at timestamp default current_timestamp
	)`

	_, err := ps.db.Exec(query)
	return err
}

func (ps *PostgresStorage) CreateAccount(account *Account) error{
	query := `
	INSERT INTO accounts 
	(first_name, last_name, number, balance, created_at) 
	VALUES 
	($1, $2, $3, $4, $5) 
	RETURNING id`
	err := ps.db.QueryRow(query, account.FirstName, account.LastName, account.Number, account.Balance, account.CreatedAt).Scan(&account.ID)
	if err != nil {
		fmt.Printf("Error inserting account: %v", err)
		return err
	}
	fmt.Printf("Inserted account with ID: %d\n", account.ID)
	return nil
}