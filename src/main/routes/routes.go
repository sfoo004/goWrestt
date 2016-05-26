package routes

import (
	"net/http"
	"main/controller"
	"github.com/gorilla/mux"
)

func RouteCalls() {
	r := mux.NewRouter()
	r.HandleFunc("/news", controller.GetAllNews).Methods("GET")
	r.HandleFunc("/news/{id:[0-9]+}", controller.GetNews).Methods("GET")
	r.HandleFunc("/news", controller.PostNews).Methods("POST")
	r.HandleFunc("/news/{id:[0-9]+}", controller.PutNews).Methods("PUT")
	r.HandleFunc("/news/{id:[0-9]+}", controller.DeleteNews).Methods("DELETE")

	//r.HandleFunc("/events", controller.GetAllEvents).Methods("GET")
	//r.HandleFunc("/events/{id:[0-9]+}", controller.GetEvents).Methods("GET")
	//r.HandleFunc("/events", controller.PostEvents).Methods("POST")
	//r.HandleFunc("/events/{id:[0-9]+}", controller.PutEvents).Methods("PUT")
	//r.HandleFunc("/events/{id:[0-9]+}", controller.DeleteEvents).Methods("DELETE")
	http.ListenAndServe(":8080", r)

}
