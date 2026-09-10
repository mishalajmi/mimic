package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mishalalajmi/mimic/internal/mock"
)

func WriteResponse(w http.ResponseWriter, r mock.Response) error {
	for k, v := range r.Headers {
		w.Header().Set(k, v)
	}

	body := r.Body

	if w.Header().Get("Content-Type") == "application/json" {
		data, err := json.Marshal(body)
		if err != nil {
			return err
		}
		body = string(data)
	}

	w.WriteHeader(r.Status)
	_, err := fmt.Fprint(w, body)
	return err
}
