package login

import (
	"encoding/json"
	"net/http"
	"robotlesson/pkg/respos"
	"strconv"
)

func NewLoginHandler(router *http.ServeMux) {
	router.HandleFunc("GET /login/", getLogin())
	router.HandleFunc("POST /login/", postLogin())
	router.HandleFunc("GET /login/{id}", getLoginId())
}

func getLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var login Login
		json.NewDecoder(r.Body).Decode(&login)
		respos.JsonResponse(w, login, 201)
	}
}

func getLoginId() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "bad id", http.StatusBadRequest)
			return
		}
		respos.JsonResponse(w, intId, http.StatusOK)
	}
}

func postLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		respos.JsonResponse(w, "POST", 201)
	}
}
