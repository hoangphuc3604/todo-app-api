package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	gin_item "social-todo-list/modules/item/transport/gin"
)

// `ID` int NOT NULL AUTO_INCREMENT,
//   `title` varchar(255) NOT NULL,
//   `image` varchar(255) DEFAULT NULL,
//   `description` text,
//   `status` enum('DOING','DONE','DELETED') NOT NULL,
//   `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
//   `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  
func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/todo_list?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()

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
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080")
}
