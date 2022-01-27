package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// a model for the car and model

type Car struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Model *Model `json:"model,omitempty"`
}

type Model struct {
	Type  string `json:"type,omitempty"`
	Year  string `json:"year,omitempty"`
	Color string `json:"color,omitempty"`
}

//a varible called cars that can have the slice of the model Car
var cars []Car

// func that handles getting of all the cars endpoint
func GetCarsEndPoint(response http.ResponseWriter, request *http.Request) {
	json.NewEncoder(response).Encode(cars)

}

// func the handles getting a single car endpoint
func GetSingleCarEndPoint(response http.ResponseWriter, request *http.Request) {
	// getting the parameters of the global var cars
	params := mux.Vars(request)

	// looping through and check if the request ID is equal to any of the Id in the global var cars
	for _, item := range cars {
		// if the request Id is equal to the Id in the loop items encode it and return
		if item.ID == params["id"] {
			json.NewEncoder(response).Encode(item)
			return
		}
	}
	//if the request ID is not equal to any of the id
	json.NewEncoder(response).Encode("Car not found")
}

// func that creates a new car endpoint
func CreateNewCarEndPoint(response http.ResponseWriter, request *http.Request) {

	//declare a new var has the model car as a value
	var newCar Car

	// decode the body and store it to the newCar var
	_ = json.NewDecoder(request.Body).Decode(&newCar)

	// append the newCar to the global var Cars
	cars = append(cars, newCar)

	json.NewEncoder(response).Encode(newCar)

}

// func that update an existin car endpoint
func UpdateCarEndPoint(response http.ResponseWriter, request *http.Request) {

	var updateCar Car
	_ = json.NewDecoder(request.Body).Decode(&updateCar)

	params := mux.Vars(request)
	for i, item := range cars {
		if item.ID == params["id"] {
			cars[i] = updateCar
			json.NewEncoder(response).Encode(updateCar)
			break
		}
	}
}

// func that deletes an existing car endpoint
func DeleteCarEndPoint(response http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)

	for i, item := range cars {
		if item.ID == params["id"] {
			copy(cars[i:], cars[i+1:])
			break
		}
	}
	json.NewEncoder(response).Encode(cars)

}

func main() {

	// calling the variable cars and populating it with data
	cars = append(cars, Car{ID: "1", Name: "Toyota", Model: &Model{Type: "Avolon", Year: "2020", Color: "Blue"}})
	cars = append(cars, Car{ID: "2", Name: "Toyota", Model: &Model{Type: "Camry", Year: "2021", Color: "Black"}})
	cars = append(cars, Car{ID: "3", Name: "Toyota", Model: &Model{Type: "Sienna", Year: "2022", Color: "Red"}})

	// setting the router using gorillamux
	router := mux.NewRouter()

	router.HandleFunc("/cars", GetCarsEndPoint).Methods("GET") // routing the get all the cars with a method Get

	router.HandleFunc("/cars/{id}", GetSingleCarEndPoint).Methods("GET") // routing the get single car endpoint with a method Get and giving it an id

	router.HandleFunc("/cars", CreateNewCarEndPoint).Methods("POST") // routing the create new car endpoint with a method Post

	router.HandleFunc("/cars/{id}", UpdateCarEndPoint).Methods("PUT") // routing the update car endpoint with a method put

	router.HandleFunc("/cars/{id}", DeleteCarEndPoint).Methods("DELETE") // routing the delete car endpoint with a method post

	fmt.Println("Server running at Port 9090")
	log.Fatal(http.ListenAndServe(":9090", router)) // setting the server port to 9090,

}
