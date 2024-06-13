package main

import (
	"bookProject/db"
	_ "bookProject/db"
	"bookProject/model"
	_ "bookProject/model"
	"fmt"
	_ "fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

type Response map[string]any

func main() {

	db.Init()
	app := gin.Default()

	app.GET("/books", func(context *gin.Context) {
		fmt.Println("Books")

		result, err := model.GetAllBooks()
		if err != nil {
			context.JSON(400, Response{
				"message": "Cannot serve your request",
			})
			return
		}
		context.JSON(200, Response{
			"message": "All books in the db",
			"books":   result,
		})
	})

	app.POST("/books", func(context *gin.Context) {

		var bookObject model.Book
		err := context.ShouldBindJSON(&bookObject)
		if err != nil {
			context.JSON(400, Response{
				"message": "Invalid object",
			})
			return
		}
		// bookObject.Id = 1
		err = bookObject.Save()

		if err != nil {
			context.JSON(400, Response{
				"message": "Cannot insert book object",
			})
			return
		}

		context.JSON(200, Response{
			"message": "Book created successfuly",
			"object":  bookObject,
		})
	})

	err := app.Run(":8080")
	if err != nil {
		fmt.Println("SERVER exception")
		fmt.Println(err)
	}
}
