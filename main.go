package squad_manager

import (
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
		panic(error(err))
	}

	Routes.CreateRoutes(*r)
	http.ListenAndServe(":80", r)
}
