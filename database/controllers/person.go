package controllers

import (
	"github.com/adrian-morfi/grapql-go-api/database/storage"
	"github.com/adrian-morfi/grapql-go-api/graphql/models"
)

func GetPersons() ([]models.Persons, error) {
	var resultPersons []models.Persons

	persons, err := GetRepoPersons()

	if err != nil {
		return nil, err
	}

	return resultPersons, nil
}

func GetRepoPersons() ([]models.Persons, error) {
	db := storage.GetDBInstance()
	persons := []models.Persons{}

	if err := db.Find(&persons).Error; err != nil {
		return nil, err
	}

	return persons, nil
}
