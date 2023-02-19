package transport

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gen95mis/short-url/internal/model"
	"github.com/gen95mis/short-url/internal/service"
)

func Service(s *service.Service) error {
	http.HandleFunc("/short", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		url := new(model.URL)
		if err := json.NewDecoder(r.Body).Decode(url); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		hash, err := s.Set(url.Original)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		shortenedUrl := fmt.Sprintf("http://localhost:/%s", hash)

		response, err := json.Marshal(shortenedUrl)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(response)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		oridinal, err := s.Get(r.URL.Path)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, oridinal, http.StatusFound)
	})

	return http.ListenAndServe(":80", nil)
}
