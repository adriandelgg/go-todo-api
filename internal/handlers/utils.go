package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"
)

func parseBodyToJSON[T any](body io.ReadCloser) (bodyParsed T, err error) {
	err = json.NewDecoder(body).Decode(&bodyParsed)
	return
}

func writeToJSON(w http.ResponseWriter, val any) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(val); err != nil {
		log.Println(err)
		http.Error(w, "Failed to send as JSON.", http.StatusInternalServerError)
	}
}

func getIdParam(w http.ResponseWriter, r *http.Request) uint {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		if errors.Is(err, strconv.ErrSyntax) {
			http.Error(w, "Invalid ID given to route parameter.", http.StatusBadRequest)
		} else {
			http.Error(w, "Error converting string to number.", http.StatusBadRequest)
		}
		return 0
	}

	return uint(id)
}
