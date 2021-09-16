package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	todoEntity "poc-svelte-golang/back/src/entity"
	todoService "poc-svelte-golang/back/src/service"

	"github.com/gorilla/mux"
)

func getTodosHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(todoService.GetMany())
}

func postTodoHandler(w http.ResponseWriter, r *http.Request) {
	newTodo := todoEntity.MakeTodo()

	if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
			fmt.Println(err)
			http.Error(w, "Error decoidng response object", http.StatusBadRequest)
			return
	}

	todoService.Add(newTodo) 
	
	json.NewEncoder(w).Encode(todoService.GetMany())
}

func putTodoChecked(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	
	todoUpdated := todoService.Check(id)

	json.NewEncoder(w).Encode(todoUpdated)
}

func TodoRouter(r *mux.Router) {
	todoRouter := r.PathPrefix("/todo").Subrouter()

	todoRouter.HandleFunc("/", getTodosHandler).Methods(http.MethodGet)
	todoRouter.HandleFunc("/", postTodoHandler).Methods(http.MethodPost)
	todoRouter.HandleFunc("/{id}", putTodoChecked).Methods(http.MethodPut)
}