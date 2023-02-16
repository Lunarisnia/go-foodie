package food

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lunarisnia/go-foodie/models"
	"github.com/lunarisnia/go-foodie/utils"
)

// TODO: Move logic to its own service
type FoodHandler struct {
	FavoriteFoodHandler *FavoriteFoodHandler
}

func (h *FoodHandler) handleGet(foodId string, res http.ResponseWriter, req *http.Request) {
	food := models.FoodModel.FindOneById(foodId)
	if food.ID == 0 {
		http.NotFound(res, req)
	}
	// TODO: add error handling
	a, _ := json.Marshal(food)
	res.Write(a)
}

func (h *FoodHandler) handlePost(res http.ResponseWriter, req *http.Request) {
	var body models.Food
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	newFood := models.FoodModel.Create(body.Name, body.Price)

	res.Write([]byte(fmt.Sprintf(`{"NewFoodID": %v }`, newFood)))
}

func (h *FoodHandler) handlePut(foodId string, res http.ResponseWriter, req *http.Request) {
	if foodId == "" {
		http.Error(res, "invalid food id", http.StatusBadRequest)
		return
	}

	var body models.Food
	jsonErr := json.NewDecoder(req.Body).Decode(&body)
	if jsonErr != nil {
		http.Error(res, jsonErr.Error(), http.StatusBadRequest)
		return
	}
	food := models.Food{Name: body.Name, Price: body.Price}
	models.FoodModel.Updates(food, "id = ?", foodId)

	fmt.Fprint(res, "Update success!")
}

func (h *FoodHandler) handleDelete(foodId string, res http.ResponseWriter, req *http.Request) {
	if foodId == "" {
		http.Error(res, "invalid food id", http.StatusBadRequest)
		return
	}

	food := models.FoodModel.FindOneById(foodId)
	models.FoodModel.Delete(food)

	fmt.Fprint(res, "Delete success!")
}

func (h *FoodHandler) Handler() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("content-type", "application/json")

		var head string
		head, req.URL.Path = utils.ShiftPath(req.URL.Path)
		switch head {
		case "favorite":
			h.FavoriteFoodHandler.Handler().ServeHTTP(res, req)
			return
		}

		if req.URL.Path == "/" {
			foodId := head
			switch req.Method {
			case "GET":
				h.handleGet(foodId, res, req)
				return
			case "POST":
				h.handlePost(res, req)
				return
			case "PUT":
				h.handlePut(foodId, res, req)
				return
			case "DELETE":
				h.handleDelete(foodId, res, req)
				return
			}
		}
		http.NotFound(res, req)
	})
}
