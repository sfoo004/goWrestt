package controller

import (
	"net/http"
	"github.com/gorilla/mux"

	"io/ioutil"
)

func GetAllNews(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello world"))
	//send request to database

	//retrieve results from DB. Parse to JSON to send back
}

func GetNews(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	w.Write([]byte(vars["id"]))
}
func PostNews(w http.ResponseWriter, req *http.Request){
	w.Write([]byte("Welcome to HELLo"))
	err := req.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//validate input
	//send request to database

	//retireve results from DB. Parse to JSON to send back
}
func PutNews(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	w.Write([]byte(vars["id"]))
	err := req.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//validate input
	//validate input id matches with id
	//send request to database

	//retrieve results from DB. Parse to JSON to send back
}

func DeleteNews(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	w.Write([]byte(vars["id"]))
	//send request to database

	//retrieve results from DB. Parse to JSON to send back
}
