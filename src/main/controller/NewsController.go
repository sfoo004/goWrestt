package controller

import (
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"encoding/json"
	"log"
	"time"
	"strconv"
	"os"
	"math/rand"
	"mime/multipart"
	"io"
)

type news struct {
	Id int `json:"id"`
	CreatedAt time.Time`json:"createdAt"`
	Title string `json:"title"`
	Body string `json:"body"`
	Image http.File `json:"image"`
}

func GetAllNews(res http.ResponseWriter, req *http.Request) {
	//send request to database
	n := news{}
	db, e := sql.Open("mysql", "root:root@tcp(localhost:3306)/news_DB?parseTime=true")
	if( e != nil){
		fmt.Print(e)
	}

	//set mime type to JSON
	res.Header().Set("Content-type", "application/json")

	//can't define dynamic slice in golang
	newsList := [] news{};

	response, err := db.Query("SELECT * FROM news")
	if err != nil {
		fmt.Print( err )
	}
	// Close closes the Rows, preventing further enumeration. If Next returns false, the Rows are closed automatically
	defer response.Close()

	for response.Next() {
		err = response.Scan( &n.Id, &n.CreatedAt, &n.Title, &n.Body )
		if err != nil {
			fmt.Println(err)
			return
		}
		newsList = append(newsList, n)
	}

	err = response.Err()
	if err != nil {
		log.Fatal(err)
	}
	outgoingJSON, _ := json.Marshal(newsList)
	res.WriteHeader(http.StatusOK)
	fmt.Fprintln(res, string(outgoingJSON))

	//retrieve results from DB. Parse to JSON to send back
}

func GetNews(res http.ResponseWriter, req *http.Request) {
	//set mime type to JSON
	res.Header().Set("Content-type", "application/json")
	n := news {}
	var id int

	//send request to database
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/news_DB?parseTime=true")
	if( err != nil){
		fmt.Print(err)
	}

	vars := mux.Vars(req)
	//gets id from url
	id, err = strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Print(err)
	}


	if err != nil {
		http.Error(res, fmt.Sprintf("error parsing url %v", err), 500)
	}
	// will Query the news and find it by the id
	err = db.QueryRow("SELECT * FROM news WHERE id= ? ", id).Scan(&n.Id, &n.CreatedAt, &n.Title, &n.Body)
	if err != nil{
		fmt.Print( err )
	}
	outgoingJSON, _ := json.Marshal(n)
	res.WriteHeader(http.StatusOK)
	fmt.Fprintln(res, string(outgoingJSON))
}

func PostNews(res http.ResponseWriter, req *http.Request){
	//set mime type to JSON
	res.Header().Set("Content-type", "application/json")

	err := req.ParseForm()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	var n news
	//err = json.NewDecoder(req.Body).Decode(&n)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
	n.CreatedAt = time.Now()
	//send request to database
	_, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/news_DB?parseTime=true")
	if( err != nil){
		fmt.Print(err)
	}
	// the FormFile function takes in the POST input id file
	file, _, err := req.FormFile("file")

	if err != nil {
		fmt.Fprintln(res, err)
		return
	}
	defer file.Close()
	processFile(file)

	//st, err := db.Prepare("INSERT INTO news(createdAt, title, body) VALUES(?, ?, ?)")
	//if err != nil{
	//	fmt.Print( err )
	//}
	//exe, err := st.Exec(n.CreatedAt, n.Title, n.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//lastId, err := exe.LastInsertId()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//n.Id = int (lastId)
	//
	//err = db.QueryRow("SELECT * FROM news WHERE id= ? ", n.Id).Scan(&n.Id, &n.CreatedAt, &n.Title, &n.Body)
	//if err != nil{
	//	fmt.Print( err )
	//}
	//outgoingJSON, _ := json.Marshal(n)
	//res.WriteHeader(http.StatusCreated)
	//fmt.Fprintln(res, string(outgoingJSON))
	//validate input
	//retireve results from DB. Parse to JSON to send back
}

func PutNews(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	//gets id from url
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Print(err)
	}
	err = req.ParseForm()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	n := news {}
	err = json.NewDecoder(req.Body).Decode(&n)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
	n.CreatedAt = time.Now()
	//send request to database
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/news_DB?parseTime=true")
	if( err != nil){
		fmt.Print(err)
	}

	st, err := db.Prepare("UPDATE news SET createdAt = ?, title = ?, body = ? WHERE id = ?")
	if err != nil{
		fmt.Print( err );
	}
	_, err = st.Exec( n.CreatedAt, n.Title, n.Body, id)
	if err != nil {
		log.Fatal(err)
	}
	err = db.QueryRow("SELECT * FROM news WHERE id= ? ", id).Scan(&n.Id ,&n.CreatedAt, &n.Title, &n.Body)
	if err != nil{
		fmt.Print( err )
	}
	outgoingJSON, _ := json.Marshal(n)
	res.WriteHeader(http.StatusOK)
	fmt.Fprintln(res, string(outgoingJSON))
	//validate input

	//retrieve results from DB. Parse to JSON to send back
}

func DeleteNews(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	//gets id from url
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Print(err)
	}
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/news_DB?parseTime=true")
	if( err != nil){
		fmt.Print(err)
	}

	st, err := db.Prepare("DELETE FROM news WHERE id = ?")
	if err != nil{
		fmt.Print( err );
	}
	exe, err := st.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := exe.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	res.WriteHeader(http.StatusOK)
	fmt.Println(lastId)

	//retrieve results from DB. Parse to JSON to send back
}

func processFile(file multipart.File){
	log.Print(file)
	dir, err := os.Getwd()
	if(err != nil){

	}
	log.Print(dir)
	err = os.MkdirAll(dir + "/uploads/images", 0775)
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, 20)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	out, err := os.Create("/uploads/images/"+ string(b))
	if err != nil {
		log.Print(err)
		return
	}

	defer out.Close()

	// write the content from POST to the file
	_, err = io.Copy(out, file)
	if err != nil {
		log.Print(err)
	}

	log.Print("File uploaded successfully : ")
		//log.Print(string(b))
	//fi, err := file.Stat()
	//if(err != nil){
	//	log.Print("ERROR")
	//}
	//log.Print(fi.Size())
	//log.Print(fi.Name())
	//os.Rename(fi.Name(), dir + "/uploads/images/"+ string(b))


}

///**
//     * Processes an uploaded file and returns its FQDN path
//     *
//     * @param  UploadedFile $file
//     * @return string       $fileFqdn
//     */
//protected function processFileUpload(UploadedFile $file)
//{
//$currentUserId = $this->get('security.token_storage')
//->getToken()
//->getUser()
//->getId();
//
//// Prepare file destination directory path
//$appRootDir = str_replace('/app', '', $this->get('kernel')->getRootDir());
//$uploadsDir = str_replace('{id}', $currentUserId, $this->container->getParameter('learning_object_bundle.user.uploads.directory'));
//$destDir = $appRootDir.$uploadsDir;
//
//// Create destination directory if it does not exist
//if (!is_dir($destDir)) {
//mkdir($destDir, 0775, true);
//}
//
//// Get randomized file name
//$fileExtension = !$file->guessExtension() ? 'bin' : $file->guessExtension();
//$filename = sha1(uniqid(mt_rand(), true)).'.'.$fileExtension;
//
//// Move file to permanent location
//$file->move($destDir, $filename);
//
//// Prepare FQDN web path to file
//if (strpos($uploadsDir, '/web') === false) {
//$fileFqdn = "FQDN is not available. File was saved in non-public directory.";
//} else {
//$fileFqdn = $this->getParameter('app_url').str_replace('/web', '', $uploadsDir).'/'.$filename;
//}
//
//return $fileFqdn;
//}
