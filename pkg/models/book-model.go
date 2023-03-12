package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/muhammadfassihhaider/go-sql-book-management-system/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name      string `gorm:""json:"name"`
	Publisher string `json:"publisher"`
	Author    string `json:"author"`
}

func init() {
	config.ConnectToDB()
	db := config.GetDB()
	db.AutoMigrate(&Book{})
}

func GetAllBooks() []Book {
	fmt.Println("GetAllBooks")
	var Books []Book
	db.Find(&Books) // can I not pass it with & because slices are already passed by reference?
	return Books
}

func GetBookById(id uint64) (*Book, *gorm.DB) {
	var Book Book
	db := db.Where("ID=?", id).Find(&Book)
	return &Book, db
}

func (book *Book) CreateBook() *Book {
	db.NewRecord(book)
	// db.Create(&book) // what is the difference b/n this and db.Create(*book)
	return book
}

func DeleteBookById(id uint64) *Book {
	var book Book
	db.Where("ID=?", id).Delete(&book)
	return &book
}
