package component

import (
	"encoding/json"
	"net/http"
)

func Input(w http.ResponseWriter, r *http.Request) {

	err := json.NewDecoder(r.Body).Decode(&r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}
