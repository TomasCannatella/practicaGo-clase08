package main

import (
	"errors"
	"fmt"
)

const ErrorClientExist = "Error: client already exists"

var InvalidClientData error = errors.New("Invalid Data")

type Client struct {
	File,
	Name string
	Id,
	PhoneNumber int
	Home string
}

type ClientRepository struct {
	Clients []Client
}

func NewClientRepository() *ClientRepository {
	return &ClientRepository{
		Clients: make([]Client, 0),
	}
}

func (c *ClientRepository) SaveClient(newClient Client) bool {
	if check, err := checkClientData(newClient); check != true {
		if errors.Is(err, InvalidClientData) {
			fmt.Println(err)
		} else {
			fmt.Println("unexpected error occurred")
		}

		return false
	}
	for _, client := range c.Clients {
		if client.Id == newClient.Id {
			panic(ErrorClientExist)
		}
	}

	return true
}

func checkClientData(newClient Client) (bool, error) {
	if newClient.File == "" {
		return false, fmt.Errorf("%w. Client File is null", InvalidClientData)
	}
	if newClient.Name == "" {
		return false, fmt.Errorf("%w. Client Name is null", InvalidClientData)
	}
	if newClient.Id == 0 {
		return false, fmt.Errorf("%w. Client Id is null", InvalidClientData)
	}
	if newClient.PhoneNumber == 0 {
		return false, fmt.Errorf("%w. Client PhoneNum is null", InvalidClientData)
	}
	if newClient.Home == "" {
		return false, fmt.Errorf("%w. Client Home is null", InvalidClientData)
	}

	return true, nil
}

func main() {
	defer func() {
		fmt.Println("End of execution")
		if r := recover(); r != nil {
			fmt.Println("Several errors were detected at runtime")
		}
	}()

	repositoryClient := NewClientRepository()

	client1 := Client{
		File:        "001",
		Name:        "Marcos",
		Id:          1,
		PhoneNumber: 123456,
		Home:        "Perez 455",
	}

	client2 := Client{
		File:        "002",
		Name:        "Matias",
		Id:          2,
		PhoneNumber: 156789,
		Home:        "Perez 455",
	}

	repositoryClient.Clients = append(repositoryClient.Clients, client1, client2)

	clientInvalidData := Client{
		File:        "003",
		Name:        "",
		Id:          1,
		PhoneNumber: 156789,
		Home:        "Perez 455",
	}

	repositoryClient.SaveClient(clientInvalidData)

	repeatedClient := client2
	repositoryClient.SaveClient(repeatedClient)

}
