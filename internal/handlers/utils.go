package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func decodeBodyToJSON[T any](body io.ReadCloser) (bodyParsed T, err error) {
	err = json.NewDecoder(body).Decode(&bodyParsed)
	return
}

func writeJSONResponse[T any](w http.ResponseWriter, val T) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(val)
}

func deferBodyClose(Body io.ReadCloser) {
	if err := Body.Close(); err != nil {
		log.Printf("Error closing request body: %v", err)
	}
}

func getIdParam(r *http.Request) (uint, error) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		if errors.Is(err, strconv.ErrSyntax) {
			return 0, fmt.Errorf("invalid ID given to route parameter: %v", err)
		}
		return 0, fmt.Errorf("error converting string to number: %v", err)
	}

	return uint(id), nil
}
