package Controllers

import (
	"encoding/json"
	"net/http"
	"squad-manager/Models"
)

func CreateDev(w http.ResponseWriter, r *http.Request) {

	dev := &Models.Dev{}
	json.NewDecoder(r.Body).Decode(&dev)
	err := Models.InsertDev(dev)
	if err != nil {
		panic(error(err))
	}

}

func IndexDevs(w http.ResponseWriter, r *http.Request) {
	devs, err := Models.SearchAllDevs()
	if err != nil {
		panic(error(err))
	}
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