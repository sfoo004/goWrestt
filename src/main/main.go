package main

import (
	"net/http"
	"main/routes"

	//"encoding/json"
	//"main/controller"
	"github.com/gorilla/mux"
)

//type stuff []news

func main() {
	routes.RouteCalls()
}




func getAllNews(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello world"))
}

func getNews(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	w.Write([]byte(vars["id"]))
}
func postNews(w http.ResponseWriter, req *http.Request){
	w.Write([]byte("Welcome to HELLo"))
}
func putNews(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	w.Write([]byte(vars["id"]))
}

func deleteNews(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	w.Write([]byte(vars["id"]))
}
