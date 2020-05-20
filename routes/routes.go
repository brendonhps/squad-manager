package routes

import (
	"fmt"
	"net/http"
	"squad-manager/controllers"

	"github.com/gorilla/mux"
)

// CreateRoutes is a function to create routes
func CreateRoutes(router *mux.Router) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("pong")
	})
	router.HandleFunc("/dev", controllers.CreateDev).Methods("POST")
	router.HandleFunc("/dev", controllers.IndexDevs).Methods("GET")
	//router.HandleFunc("/dev/{id}", controllers.DeleteDev).Methods("DELETE")
	//router.HandleFunc("/dev/{id}", controllers.UpdateDev).Methods("PUT")

	router.HandleFunc("/project", controllers.CreateProject).Methods("POST")
	router.HandleFunc("/project", controllers.IndexProjects).Methods("GET")
	//router.HandleFunc("/project/{id}", controllers.DeleteProject).Methods("DELETE")
	//router.HandleFunc("/project/{id}", controllers.UpdateProject).Methods("UPDATE")

	router.HandleFunc("/project", controllers.CreateDeploy).Methods("POST")
	router.HandleFunc("/project", controllers.IndexDeploys).Methods("GET")
	//router.HandleFunc("/project/{id}", controllers.DeleteDeploy).Methods("DELETE")
}
