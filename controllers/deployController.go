package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"squad-manager/models"
)

//CreateDeploy is a function to handle a route to create a deploy
func CreateDeploy(w http.ResponseWriter, r *http.Request) {
	deploy := models.Deploy{}
	err := json.NewDecoder(r.Body).Decode(&deploy)
	if err != nil {
		log.Fatal(err)
	}

	// Todo validate requirement id

	err = Models.InsertDeploy(deploy)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(200)
	err = json.NewEncoder(w).Encode(deploy)
	if err != nil {
		panic(error(err))
	}
}

//IndexDeploys is a function to handle a route to get all deploys 
func IndexDeploys(w http.ResponseWriter, r *http.Request) {

}

/*TODO*/
//func DeleteDeploy(w http.ResponseWriter, r *http.Request) {
//
//}
