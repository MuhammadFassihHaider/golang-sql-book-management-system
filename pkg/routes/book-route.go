package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/muhammadfassihhaider/go-sql-book-management-system/pkg/controllers"
)

func RegisterBookRoutes(router *mux.Router) {
	router.HandleFunc("/books/", controllers.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/", controllers.CreateBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{bookId}", controllers.GetBookById).Methods(http.MethodGet)
	router.HandleFunc("/books/{bookId}", controllers.DeleteBookById).Methods(http.MethodDelete)
	router.HandleFunc("/books/{bookId}", controllers.UpdateBookById).Methods(http.MethodPut)
}
