package Controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"squad-manager/Models"
)

func CreateDev(w http.ResponseWriter, r *http.Request) {
	dev := Models.Dev{}
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

func IndexDevs(w http.ResponseWriter, r *http.Request) {
	devs, err := Models.SearchAllDevs()
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
