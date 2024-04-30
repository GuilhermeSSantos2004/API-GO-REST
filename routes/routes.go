package routes

import (
	"net/http"

	"github.com/GuilhermeSSantos2004/API-GO-REST/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.http.HandleFunc("/", controllers.Home)
	r.http.HandleFunc("/api/personalidades", controllers.TodasPersonalidades)
	r.log.Fatal(http.ListenAndServe(":8000", r))
}
