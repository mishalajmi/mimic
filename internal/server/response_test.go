package server

import (
	"net/http/httptest"
	"testing"

	"github.com/mishalalajmi/mimic/internal/mock"
)

func TestWriteJSONResponse(t *testing.T) {
	response := mock.Response{
		Status: 201,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: map[string]any{
			"id":     123,
			"status": "approved",
		},
	}

	recorder := httptest.NewRecorder()

	err := WriteResponse(recorder, response)
	if err != nil {
		t.Fatal(err)
	}

	if recorder.Code != 201 {
		t.Fatalf("expected status 201, got %d", recorder.Code)
	}

	if recorder.Header().Get("Content-Type") != "application/json" {
		t.Fatal("expected JSON content type")
	}

	expected := `{"id":123,"status":"approved"}`

	if recorder.Body.String() != expected {
		t.Fatalf("expected %s, got %s", expected, recorder.Body.String())
	}
}
