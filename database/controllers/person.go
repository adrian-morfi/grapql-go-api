package controllers

import (
	"api/database/storage"
	"api/graphql/models"
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
