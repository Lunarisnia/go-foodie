package main

import (
	"log"
	"net/http"

	"github.com/lunarisnia/go-foodie/api/routers"
	"github.com/lunarisnia/go-foodie/api/routers/food"
	"github.com/lunarisnia/go-foodie/models"
)

func main() {
	// Connect to the database
	models.ConnectDatabase()

	a := &routers.App{
		FoodHandler: new(food.FoodHandler),
	}
	err := http.ListenAndServe(":3000", a)
	if err != nil {
		log.Fatal(err)
	}
}
