package main

import (
	"api/database/storage"
	"api/graphql/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//db := models.SetupModels()
	storage.NewDB()

	defer db.Close()

	/*r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	*/

	r.GET("/graphql", controllers.Graphql)

	r.Run()
}
