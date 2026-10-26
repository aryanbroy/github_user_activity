package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func FetchUserActivity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	fmt.Println("Username: ", username)
}
