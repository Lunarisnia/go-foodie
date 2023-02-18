package addOn

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lunarisnia/go-foodie/models"
	"github.com/lunarisnia/go-foodie/services/addOnServices"
	"github.com/lunarisnia/go-foodie/utils"
)

type AddOnHandler struct {
}

// TODO: Split logic further
func (h *AddOnHandler) handleGet(id string, res http.ResponseWriter, req *http.Request) {
	if id == "" {
		http.NotFound(res, req)
		return
	}

	addOn := models.AddOnModel.FindOneById(id)
	if addOn.ID == 0 {
		http.Error(res, "add on not found", http.StatusNotFound)
		return
	}

	addOnJson, err := json.Marshal(addOn)
	if err != nil {
		http.Error(res, "internal server error", http.StatusInternalServerError)
		return
	}
	res.Write(addOnJson)
}

func (h *AddOnHandler) handlePost(res http.ResponseWriter, req *http.Request) {
	var body models.AddOn
	eofErr := utils.BodyToJson(req.Body, &body)
	if eofErr != nil {
		http.Error(res, "bad request", http.StatusBadRequest)
		return
	}

	response, err := addOnServices.AddNewAddOn(body)
	if err != nil {
		http.Error(res, "bad request", http.StatusBadRequest)
		return
	}
	res.Write(response)
}

func (h *AddOnHandler) handlePut(addOnId string, res http.ResponseWriter, req *http.Request) {
	if addOnId == "" {
		http.Error(res, "invalid add on id", http.StatusBadRequest)
		return
	}

	var body models.AddOn
	eofErr := utils.BodyToJson(req.Body, &body)
	if eofErr != nil {
		http.Error(res, "bad request", http.StatusBadRequest)
		return
	}

	addOn := models.AddOn{Name: body.Name, Price: body.Price, PhotoUrl: body.PhotoUrl, InStock: body.InStock, FoodID: body.FoodID}
	models.AddOnModel.Updates(addOn, "id = ?", addOnId)

	fmt.Fprint(res, "Update success!")
}

func (h *AddOnHandler) handleDelete(addOnId string, res http.ResponseWriter, req *http.Request) {
	if addOnId == "" {
		http.Error(res, "invalid add on id", http.StatusBadRequest)
		return
	}

	addOn := models.AddOnModel.FindOneById(addOnId)
	if addOn.ID == 0 {
		http.Error(res, "add on not found", http.StatusNotFound)
		return
	}
	models.AddOnModel.Delete(addOn)

	fmt.Fprint(res, "Delete success!")
}

func (h *AddOnHandler) Handler() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("content-type", "application/json")

		var head string
		head, req.URL.Path = utils.ShiftPath(req.URL.Path)
		switch head {
		default:
			break
		}

		if req.URL.Path == "/" {
			id := head
			switch req.Method {
			case "GET":
				h.handleGet(id, res, req)
				return
			case "POST":
				h.handlePost(res, req)
				return
			case "PUT":
				h.handlePut(id, res, req)
				return
			case "DELETE":
				h.handleDelete(id, res, req)
				return
			}
		}
		http.NotFound(res, req)
	})
}
