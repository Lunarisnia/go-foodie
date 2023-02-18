package addOnServices

import (
	"fmt"

	"github.com/lunarisnia/go-foodie/models"
	"github.com/lunarisnia/go-foodie/services/types"
)

func AddNewAddOn(body models.AddOn) ([]byte, error) {
	newAddOn := models.AddOnModel.Create(body.Name, body.PhotoUrl, body.Price, body.InStock, body.FoodID)
	if newAddOn == 0 {
		return nil, &types.FailedCreateData{}
	}
	response := []byte(fmt.Sprintf(`{"NewAddOnID": %v }`, newAddOn))
	return response, nil
}
