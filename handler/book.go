package handler

import (
	"github.com/gin-gonic/gin"
	"golang-web-testcase/models"
	"golang-web-testcase/utility"
	"net/http"
	"strconv"
)

//ShowIndexPage show index
func ShowIndexPage(c *gin.Context) {
	books := models.GetAllBooks()
	utility.Render(
		c,
		gin.H{
			"title":   "Hi",
			"payload": books,
		},
		"index.html",
	)
}

//GetBook get book
func GetBook(c *gin.Context) {
	if bookID, err := strconv.Atoi(c.Param("book_id")); err == nil {
		book := models.GetBookByID(bookID)
		utility.Render(
			c,
			gin.H{
				"title":   "Book",
				"payload": book,
			},
			"book.html",
		)
	} else {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}

//SaveBook save book
func SaveBook(c *gin.Context) {
	var book models.Book
	if err := c.Bind(&book); err == nil {
		book := models.SaveBook(book)
		utility.Render(
			c,
			gin.H{
				"title":   "Save",
				"payload": book,
			},
			"success.html",
		)
	} else {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}

//ShowSaveBook show save book
func ShowSaveBook(c *gin.Context) {
	utility.Render(
		c,
		gin.H{
			"title": "SaveBook",
		},
		"add_book.html",
	)
}

//EditBook edit book
func EditBook(c *gin.Context) {
	if bookID, err := strconv.Atoi(c.Param("book_id")); err == nil {
		book := models.GetBookByID(bookID)
		utility.Render(
			c,
			gin.H{
				"title":   "Book",
				"payload": book,
			},
			"edit_book.html",
		)
	} else {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}

//EditSaveBook edit save book
func EditSaveBook(c *gin.Context) {
	var book models.Book
	if bookID, err := strconv.Atoi(c.Param("book_id")); err == nil {
		if err := c.Bind(&book); err == nil {
			book := models.EditBook(bookID, book)
			utility.Render(
				c,
				gin.H{
					"title":   "Save",
					"payload": book,
				},
				"success.html",
			)
		} else {
			c.AbortWithError(http.StatusBadRequest, err)
		}
	}
}

//DeleteBook delete book
func DeleteBook(c *gin.Context) {
	var book models.Book
	if bookID, err := strconv.Atoi(c.Param("book_id")); err == nil {
		if err := c.Bind(&book); err == nil {
			res := models.DeleteBook(bookID)
			utility.Render(
				c,
				gin.H{
					"title":  "Delete",
					"delete": res,
				},
				"success.html",
			)
		} else {
			c.AbortWithError(http.StatusBadRequest, err)
		}
	}
}
