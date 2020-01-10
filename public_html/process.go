package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"time"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


type PageVariables struct {
	Date         string
	Time         string
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", MainPage)
	r.HandleFunc("/complete", CompleteTask)
	r.HandleFunc("/create", CreateTask)
	r.HandleFunc("/delete", DeleteTask)
	r.HandleFunc("/get", GetTasks)
	r.HandleFunc("/save", SaveTask)
	log.Print("starting up")
	http.ListenAndServe(":8000", r)
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("main.html")
	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	t.Execute(w, nil)

}

func CompleteTask(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		task := r.FormValue("task")
		log.Print("in CompleteTask task: ",task)
		complete, err := db.Prepare("update tasks set complete=1-complete where id=?")
		if err != nil {
			log.Print("complete failed");
			return;
		}
		complete.Exec(task)
	}
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	db := dbConn()
	if r.Method == "POST" {
		task := r.FormValue("task")
		insert, err := db.Prepare("insert into tasks values (?,?,?,?,?)")
		if err != nil {
			log.Print("insert failed");
			return;
		}
		insert.Exec(nil, 1, task, 0, now.Unix())
	}
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		task := r.FormValue("task")
		log.Print("task:", task)
		delete, err := db.Prepare("delete from tasks where id=?")
		if err != nil {
			log.Print("delete failed");
			return;
		}
		delete.Exec(task)
	}
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	seldb, err := db.Query("select id, task, complete from tasks where user_id=?", 1)
	for seldb.Next() {
		var id int
		var task string
		var complete int
		err = seldb.Scan(&id, &task, &complete)
		if err != nil {
			log.Print("select failed")
			return;
		}
		fmt.Fprintf(w, "%d:%s:%d\n", id, task, complete)
	}

}

func SaveTask(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		task := r.FormValue("task")
		i := r.FormValue("i")
		log.Print("in SaveTask idtask:",task)
		save, err := db.Prepare("update tasks set task=? where id=?")
		if err != nil {
			log.Print("save failed");
			return;
		}
		save.Exec(task, i)
	}
}
func dbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "jg"
    dbPass := "jgpass"
    dbName := "jgdb"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}
