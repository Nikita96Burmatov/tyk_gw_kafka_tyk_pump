package helpers

import (
	"encoding/json"
	"github.com/TykTechnologies/tyk/header"
	"net/http"
	"slices"
)

var (
	entities = []string{"deal", "company", "contact"}
)

func IsAvaliableEntity(entity string) bool {
	return slices.Contains(entities, entity)
}

func DoJSONWrite(w http.ResponseWriter, code int, obj interface{}) {
	w.Header().Set(header.ContentType, header.ApplicationJSON)
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(obj); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
