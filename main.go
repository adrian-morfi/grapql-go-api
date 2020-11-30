package main

import (
	"github.com/adrian-morfi/grapql-go-api/database/storage"
	"github.com/adrian-morfi/grapql-go-api/graphql/controllers"

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
