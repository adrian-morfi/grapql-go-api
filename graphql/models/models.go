package models

import (
	"github.com/graphql-go/graphql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Todo struct {
	ID     int    `json:"id" gorm:"primary_key"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

var TodoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Todo",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"status": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

type Persons struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Address   string `json:"address"`
}

var PersonsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Persons",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"firstname": &graphql.Field{
			Type: graphql.String,
		},
		"lastname": &graphql.Field{
			Type: graphql.String,
		},
		"address": &graphql.Field{
			Type: graphql.String,
		},
	},
})

/*
func SetupModels() *gorm.DB {
	db, err := gorm.Open("postgres", "user= dbname= sslmode=disable")

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Todo{})

	return db
}
*/
