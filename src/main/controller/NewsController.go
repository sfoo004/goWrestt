package controller

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/go-sql-driver/mysql"
	"io/ioutil"
	"database/sql"
	"fmt"
	"encoding/json"
	"strconv"
)

type news struct {
	Id int `json:"id"`
	CreatedAt float32 `json:"createdAt"`
	Title string `json:"title"`
	Body string `json:"body"`
}

func GetAllNews(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello world"))
	//send request to database
	db, e := sql.Open("mysql", "username:password@tcp(localhost:3306)/news")
	if( e != nil){
		fmt.Print(e)
	}

	//set mime type to JSON
	res.Header().Set("Content-type", "application/json")

	err := req.ParseForm()
	if err != nil {
		http.Error(res, fmt.Sprintf("error parsing url %v", err), 500)
	}

	//can't define dynamic slice in golang
	var result = make([]string,1000)

	st, err := db.Prepare("SELECT * FROM news")
	if err != nil{
		fmt.Print( err );
	}
	rows, err := st.Query()
	if err != nil {
		fmt.Print( err )
	}
	i := 0
	for rows.Next() {
		var id int
		var createdAt float32
		var title string
		var body string
		err = rows.Scan( &id, &createdAt, &title, &body )
		new := news{Id:id, CreatedAt:createdAt, Title:title, Body:body}
		b, err := json.Marshal(new)
		if err != nil {
			fmt.Println(err)
			return
		}
		result[i] = fmt.Sprintf("%s", string(b))
		i++
	}
	result = result[:i]

	json, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send the text diagnostics to the client.
	fmt.Fprintf(res,"%v",string(json))
	//fmt.Fprintf(response, " request.URL.Path   '%v'\n", request.Method)
	db.Close()

	//retrieve results from DB. Parse to JSON to send back
}

func GetNews(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	res.Write([]byte(vars["id"]))
	//send request to database
	db, e := sql.Open("mysql", "username:password@tcp(localhost:3306)/news")
	if( e != nil){
		fmt.Print(e)
	}

	//set mime type to JSON
	res.Header().Set("Content-type", "application/json")

	err := req.ParseForm()
	if err != nil {
		http.Error(res, fmt.Sprintf("error parsing url %v", err), 500)
	}

	//can't define dynamic slice in golang
	var result = make([]string,1000)

	st, err := db.Prepare("SELECT * FROM news WHERE id= ? ")
	if err != nil{
		fmt.Print( err );
	}
	rows, err := st.Query()
	if err != nil {
		fmt.Print( err )
	}
	i := 0
	for rows.Next() {
		var id int
		var createdAt float32
		var title string
		var body string
		err = rows.Scan( &id, &createdAt, &title, &body )
		new := news{Id:id, CreatedAt:createdAt, Title:title, Body:body}
		b, err := json.Marshal(new)
		if err != nil {
			fmt.Println(err)
			return
		}
		result[i] = fmt.Sprintf("%s", string(b))
		i++
	}
	result = result[:i]

	json, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send the text diagnostics to the client.
	fmt.Fprintf(res,"%v",string(json))
	//fmt.Fprintf(response, " request.URL.Path   '%v'\n", request.Method)
	db.Close()
}

func PostNews(res http.ResponseWriter, req *http.Request){
	res.Write([]byte("Welcome to HELLo"))
	err := req.ParseForm()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	var n news
	decoder := json.NewDecoder(req.Body).Decode(&n)


	//validate input
	//send request to database

	//retireve results from DB. Parse to JSON to send back
}
func PutNews(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	res.Write([]byte(vars["id"]))
	err := req.ParseForm()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	//validate input
	//validate input id matches with id
	//send request to database

	//retrieve results from DB. Parse to JSON to send back
}

func DeleteNews(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	res.Write([]byte(vars["id"]))
	//send request to database

	//retrieve results from DB. Parse to JSON to send back
}
