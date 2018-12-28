package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/azbshiri/hex-rest-api/pkg/auth"
)

func NewRouter(a auth.Service) http.Handler {
	router := mux.NewRouter()
	router.POST("/join", addUser(a))
	return router
}


func addUser(s auth.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var newUser auth.User
		err := json.NewDecoder(r.Body).Decode(&newBeer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		s.AddUser(newUser)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"msg": "New user added"})
	}
}

