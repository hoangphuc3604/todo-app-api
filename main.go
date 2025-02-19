package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"social-todo-list/middleware"
	gin_item "social-todo-list/modules/item/transport/gin"
)
  
func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/todo_list?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()
	r.Use(middleware.Recovery())

	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.GET("", gin_item.ListItems(db))
			items.POST("", gin_item.CreatItem(db))
			items.GET("/:id", gin_item.GetItem(db))
			items.PATCH("/:id", gin_item.UpdateItem(db))
			items.DELETE("/:id", gin_item.DeleteItem(db))
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		fmt.Println([]int{}[0])

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080")
}
