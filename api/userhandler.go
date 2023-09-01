package api

import (
	"encoding/json"
	"littleapi/types"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *Server) handleGetUserByID(w http.ResponseWriter, r *http.Request) {
	i := getRequestId(r)
	user := s.store.Get(i)

	if user == nil {
		errorHandler(w, r, http.StatusNotFound, "User not found")
		return
	}

	json.NewEncoder(w).Encode(user)
}

func getRequestId(r *http.Request) int {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	return i
}

func (s *Server) handleDeleteUserByID(w http.ResponseWriter, r *http.Request) {
	i := getRequestId(r)
	var resp int = s.store.Delete(i)
	if resp == 0 {
		errorHandler(w, r, http.StatusNotFound, "User not found")
		return
	}

	w.WriteHeader(http.StatusOK)
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
