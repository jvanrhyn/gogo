package api

import (
	"encoding/json"
	"littleapi/types"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *Server) handleGetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	user := s.store.Get(i)

	if user == nil {
		errorHandler(w, r, http.StatusNotFound, "User not found")
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (s *Server) handleDeleteUserByID(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var user types.User
	err := decoder.Decode(&user)
	if err != nil {
		errorHandler(w, r, http.StatusBadRequest, "User data broken")
		return
	}

	result := s.store.Add(user)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

}
