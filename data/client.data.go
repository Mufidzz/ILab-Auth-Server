package data

import (
	"../structs"
	"github.com/gin-gonic/gin"
)

//Take all available Clients from Database
func (idb *InDB) GetClientsWorker() ([]structs.Client, error) {
	var (
		data []structs.Client
		err  error
	)

	fetchResult := idb.DB.
		Find(&data)

	if fetchResult.Error != nil {
		err = fetchResult.Error
		return data, err
	}

	return data, nil
}

//Take single Client using id
func (idb *InDB) GetClientUsingBarrierIDWorker(id string) (structs.Client, error) {
	var (
		data structs.Client
		err  error
	)

	fetchResult := idb.DB.
		Where("barrier_id = ?", id).
		Last(&data)

	if fetchResult.Error != nil {
		err = fetchResult.Error
		return data, err
	}

	return data, nil
}

//Append new Client to database
func (idb *InDB) CreateClientWorker(newData structs.Client) (structs.Client, error) {
	var (
		err error
	)

	fetchResult := idb.DB.
		Create(&newData)

	if fetchResult.Error != nil {
		err = fetchResult.Error
		return newData, err
	}

	return newData, nil
}

//Change available Client in database
func (idb *InDB) UpdateClientUsingIDWorker(id string, c *gin.Context) (structs.Client, error) {
	var (
		data structs.Client
		err  error
	)

	data, err = idb.GetClientUsingBarrierIDWorker(id)
	if err != nil {
		return data, err
	}

	err = c.BindJSON(&data)
	if err != nil {
		return data, err
	}

	fetchResult := idb.DB.
		Save(&data)

	if fetchResult.Error != nil {
		err = fetchResult.Error
		return data, err
	}

	return data, nil
}

//Soft Delete available Client in database
func (idb *InDB) DeleteClientUsingIDWorker(id string) (structs.Client, error) {
	var (
		data structs.Client
		err  error
	)

	data, err = idb.GetClientUsingBarrierIDWorker(id)
	if err != nil {
		return data, err
	}

	fetchResult := idb.DB.
		Delete(&data)

	if fetchResult.Error != nil {
		err = fetchResult.Error
		return data, err
	}

	return data, nil
}
