package main

import (
	"github.com/lunarisnia/go-foodie/models"
)

func main() {
	// Connect to the database
	models.ConnectDatabase()
}

// TODO: Add REST API using net/http