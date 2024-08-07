package helper

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetBody[T any](w http.ResponseWriter, r *http.Request, data T) T {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusInternalServerError)
		return data
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			http.Error(w, "Error closing request body", http.StatusInternalServerError)
			return
		}
	}(r.Body)

	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return data
	}

	return data
}
