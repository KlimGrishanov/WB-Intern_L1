package main

import "fmt"

// Order Main Struct
type Order struct {
	id       int
	price    int
	database string
}

// Repo Adapter Interface
type Repo interface {
	GetOrder() Order
}

// Work with PostgreSQL & GetOrder

type PostgreSQL struct {
	ip   string
	port string
	pass string
	user string
}

func createPostgreSQL() *PostgreSQL {
	return &PostgreSQL{ip: "127.0.0.1", port: "5432", pass: "pass", user: "admin"}
}

func (s *PostgreSQL) GetOrder() Order {
	// Business Logic Happens Here
	return Order{id: 1, price: 127, database: "PostgreSQL"}
}

// Work with MySQL & GetOrder

type MySQL struct {
	ip   string
	port string
	pass string
	user string
}

func createMySQL() *MySQL {
	return &MySQL{ip: "127.0.0.1", port: "5432", pass: "pass", user: "admin"}
}

func (s *MySQL) GetOrder() Order {
	// Business Logic Happens Here
	return Order{id: 1, price: 127, database: "MySQL"}
}

// Unify GetOrder for any DBs with adapter

func GetOrder(repo Repo) Order {
	return repo.GetOrder()
}

func main() {

	// Creating two services
	mySQL := createMySQL()
	postgreSQL := createPostgreSQL()

	// Get Order by each database
	fmt.Println(GetOrder(mySQL))
	fmt.Println(GetOrder(postgreSQL))
}
