package main

import (
	"encoding/json"
	"net/http"

	"github.com/Wellington19/posalfa-atividade-introducao-go/domain"
	"github.com/Wellington19/posalfa-atividade-introducao-go/domain/baptized"
	"github.com/Wellington19/posalfa-atividade-introducao-go/domain/barbecue"
	"github.com/Wellington19/posalfa-atividade-introducao-go/domain/birthday"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/barbecue", calculateParty(barbecue.NewBarbecue()))
	r.Post("/baptized", calculateParty(baptized.NewBaptized()))
	r.Post("/birthday", calculateParty(birthday.NewBirthday()))
	http.ListenAndServe(":3000", r)
}

func calculateParty(s domain.Party) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var param domain.Parameters
		err := json.NewDecoder(r.Body).Decode(&param)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		ch, err := s.Calculate(param)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		j, err := ch.ToJSON()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		_, err = w.Write(j)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}
}
