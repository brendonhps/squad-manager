package Controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"squad-manager/Models"
)

func CreateDeploy(w http.ResponseWriter, r *http.Request) {
	deploy := Models.Deploy{}
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

func IndexDeploys(w http.ResponseWriter, r *http.Request) {

}

/*TODO*/
//func DeleteDeploy(w http.ResponseWriter, r *http.Request) {
//
//}
