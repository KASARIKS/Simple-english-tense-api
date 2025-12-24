package tense

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kasariks/a_nameless_project_yet/types"
)

func TestTenseServiceHandlers(t *testing.T) {
	handlers := NewHandler()

	t.Run("should correctly response with tense", func(t *testing.T) {
		tense := &types.Tense{
			Name:      "Simple Present Tense",
			Structure: "Subject + verb(v1) + s/es",
			Example:   "She goes to the shop.",
		}

		req, err := http.NewRequest(http.MethodGet, "/gettense?tense=simplepresent", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := http.NewServeMux()

		router.HandleFunc("/gettense", handlers.GetTenseHandler)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d got %d", http.StatusOK, rr.Code)
		}

		var gottenTense *types.Tense = &types.Tense{}
		if err := json.NewDecoder(rr.Body).Decode(gottenTense); err != nil {
			t.Fatal(err)
		}

		if *tense != *gottenTense {
			t.Errorf("expected %v got %v", *tense, *gottenTense)
		}
	})

	t.Run("should response with 400 code to bad request", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/gettense", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := http.NewServeMux()

		router.HandleFunc("/gettense", handlers.GetTenseHandler)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d got %d", http.StatusBadRequest, rr.Code)
		}
	})
}
