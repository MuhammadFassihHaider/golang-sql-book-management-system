package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/muhammadfassihhaider/go-sql-book-management-system/pkg/routes"
)

func catchErrors() {
	if err := recover(); err != nil {
		log.Println(err)
	}
}

func main() {
	defer catchErrors()
	router := mux.NewRouter()
	routes.RegisterBookRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8000", router))
}
