package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"squad-manager/models"
)

// CreateDev is a function to handle the route to create a dev
func CreateDev(w http.ResponseWriter, r *http.Request) {
	dev := models.Dev{}
	err := json.NewDecoder(r.Body).Decode(&dev)
	if err != nil {
		fmt.Println(err)
		panic(error(err))
	}

	err = Models.InsertDev(dev)
	if err != nil {
		panic(error(err))
	}
	w.WriteHeader(200)
	err = json.NewEncoder(w).Encode(dev)
	if err != nil {
		panic(error(err))
	}
}

// IndexDevs is a function to handle the route to get all devs on the system
func IndexDevs(w http.ResponseWriter, r *http.Request) {
	devs, err := models.SearchAllDevs()
	if err != nil {
		panic(error(err))
	}
	w.WriteHeader(200)
	err = json.NewEncoder(w).Encode(devs)
	if err != nil {
		panic(error(err))
	}
}

/*TODO*/
//func DeleteDev(w http.ResponseWriter, r *http.Request) {
//
//}
//
//func UpdateDev(w http.ResponseWriter, r *http.Request) {
//
//}
