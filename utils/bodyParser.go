package utils

import (
	"encoding/json"
	"io"
)

func BodyToJson(body io.Reader, dest interface{}) error {
	err := json.NewDecoder(body).Decode(&dest)
	return err
}
