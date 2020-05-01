package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"net/http"
	"squad-manager/Routes"
)

func main() {
	r := mux.NewRouter()
	err := godotenv.Load()
	if err != nil {
		/* TODO Implement the logger to capture all events and remove the panics */
		fmt.Println("Found err in env load: ", err)
		panic(error(err))
	}

	Routes.CreateRoutes(r)
	err = http.ListenAndServe(":8000", r)
	if err != nil {
		fmt.Println("Found err on listenAndServe:", err)
		panic(error(err))
	}
}
