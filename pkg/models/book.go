package models

import (
	"github.com/kevingdc/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string
	Author      string
	Publication string
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) Create() *Book {
	db.Create(&b)
	return b
}

func (b *Book) Update() *Book {
	db.Save(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book

	db.Find(&books)
	return books
}

func GetBookById(id int64) *Book {
	book := &Book{}
	db.First(book, id)
	return book
}

func DeleteBook(id int64) *Book {
	book := GetBookById(id)
	db.Delete(book)
	return book
}
