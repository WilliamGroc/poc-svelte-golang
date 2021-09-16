package main

//https://mj-go.in/golang/crud-rest-api-with-gorilla-mux
//https://gowebexamples.com/routes-using-gorilla-mux/

import (
	"log"
	"net/http"
	"poc-svelte-golang/back/src/controller"
	todoService "poc-svelte-golang/back/src/service"

	"github.com/gorilla/mux"
)

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.Method + " " + r.URL.Path)
			w.Header().Add("Content-Type", "application/json")
			next.ServeHTTP(w, r)
	})
}

func requestHandler() {
	r := mux.NewRouter()
	r.Use(commonMiddleware)
	controller.TodoRouter(r)
	
	log.Println("Start listening to 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	todoService.Create("hello")
	todoService.Create("world")

	requestHandler()
}