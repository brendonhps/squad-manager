package squad_manager

import (
	"github.com/gorilla/mux"
	"net/http"
	"squad-manager/Routes"
)

func main() {
	r := mux.NewRouter()
	Routes.CreateRoutes(*r)
	http.ListenAndServe(":80", r)
}
