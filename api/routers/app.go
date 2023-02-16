package routers

import (
	"fmt"
	"net/http"

	"github.com/lunarisnia/go-foodie/api/routers/food"
	"github.com/lunarisnia/go-foodie/utils"
)

type App struct {
	FoodHandler *food.FoodHandler
}

func (h *App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = utils.ShiftPath(req.URL.Path)
	switch head {
	case "":
		fmt.Fprint(res, "Welcome to Foodie API! 😉")
	case "food":
		h.FoodHandler.Handler().ServeHTTP(res, req)
	default:
		http.NotFound(res, req)
	}
}
