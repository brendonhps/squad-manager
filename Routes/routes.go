package Routes

import (
	"github.com/gorilla/mux"
	"squad-manager/Controllers"
)

func CreateRoutes(router mux.Router) {
	router.HandleFunc("/dev", Controllers.CreateDev).Methods("POST")
	router.HandleFunc("/dev", Controllers.IndexDevs).Methods("GET")
	//router.HandleFunc("/dev/{id}", Controllers.DeleteDev).Methods("DELETE")
	//router.HandleFunc("/dev/{id}", Controllers.UpdateDev).Methods("PUT")

	router.HandleFunc("/project", Controllers.CreateProject).Methods("POST")
	router.HandleFunc("/project", Controllers.IndexProjects).Methods("GET")
	//router.HandleFunc("/project/{id}", Controllers.DeleteProject).Methods("DELETE")
	//router.HandleFunc("/project/{id}", Controllers.UpdateProject).Methods("UPDATE")

	router.HandleFunc("/project", Controllers.CreateDeploy).Methods("POST")
	router.HandleFunc("/project", Controllers.IndexDeploys).Methods("GET")
	//router.HandleFunc("/project/{id}", Controllers.DeleteDeploy).Methods("DELETE")
}