package resources

import (
	"encoding/json"
	"net/http"
)

func HandleHello(w http.ResponseWriter, r *http.Request) {
	person := struct {
		Name string `json:"name", db:"pk"`
		Age  int    `json:"age"`
	}{"Tom", 12}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}
