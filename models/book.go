package models

import (
	"golang-web-testcase/utility"
)

//Book model
type Book struct {
	ID           int    `json:"id"`
	Title        string `form:"title" json:"title" gorm:"type:varchar(100);unique"`
	Author       string `form:"author" json:"author" gorm:"type:varchar(100)"`
	Introduction string `form:"introduction" gorm:"type:text"`
}

//GetAllBooks get all books
func GetAllBooks() []Book {
	var books []Book
	utility.DB().Find(&books)
	return books
}

//GetBookByID get book by id
func GetBookByID(id int) Book {
	var book Book
	utility.DB().First(&book, id)
	return book
}

//SaveBook save book by id
func SaveBook(book Book) Book {
	utility.DB().Create(&book)
	return book
}
