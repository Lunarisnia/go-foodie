package food

import (
	"fmt"
	"net/http"

	"github.com/lunarisnia/go-foodie/models"
	"github.com/lunarisnia/go-foodie/utils"
)

type FavoriteFoodHandler struct {
}

func (h *FavoriteFoodHandler) handleGet(foodId string, res http.ResponseWriter, req *http.Request) {
	food := models.FoodModel.FindOneById(foodId)
	if food.ID == 0 {
		http.Error(res, "food not found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(res, "My favorite food is: %v", food.Name)
}

func (h *FavoriteFoodHandler) Handler() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var head string
		head, req.URL.Path = utils.ShiftPath(req.URL.Path)

		if req.URL.Path == "/" {
			id := head
			if id == "" {
				http.NotFound(res, req)
				return
			}
			switch req.Method {
			case "GET":
				h.handleGet(id, res, req)
				return
			}
		}
		http.NotFound(res, req)
	})
}
