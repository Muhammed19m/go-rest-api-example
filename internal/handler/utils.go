package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func UnmarshalBody(r *http.Request, a any) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return err
	}
	defer r.Body.Close()
	err = json.Unmarshal(body, a)
	if err != nil {
		// http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return err
	}

	return nil
}