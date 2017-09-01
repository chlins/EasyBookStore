package router

import (
	"github.com/gin-gonic/gin"
	"golang-web-testcase/handler"
)

func Load() *gin.Engine {
	router := gin.Default()
	router.GET("/", handler.ShowIndexPage)
	router.GET("/book/:book_id", handler.GetBook)
	router.GET("/add_book", handler.ShowSaveBook)
	router.POST("/book", handler.SaveBook)
	router.GET("/book/:book_id/edit", handler.EditBook)
	router.POST("/editbook/:book_id", handler.EditSaveBook)
	router.GET("delete/:book_id", handler.DeleteBook)
	return router
}
