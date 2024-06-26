package handler

import (
	"encoding/json"
)

func jsonError(msg string) []byte {
	errStruct := struct {
		Message string `json:"message"`
	}{
		msg,
	}

	errJson, err := json.Marshal(errStruct)
	if err != nil {
		return []byte(err.Error())
	}

	return errJson
}
