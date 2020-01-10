package main
import (
	"io/ioutil"
	"net/http"
    "net/http/httptest"
	"net/url"
    "testing"
	"strings"
	"log"
)

// This module POSTS (creates) a task with a test entry. Then, GETs it.
// Many more tests could be added here.

const (
	test_url string = "http://localhost:8000/"
	test_task string = "TEST entry - Feed the cats and dogs and sambo and goldie - please delete"
	test_title string = "Task Keeper&trade; - POC - Jahnel Group - Korey O'Dell - Phase 2 deliverable"
)

func TestCreateTask(t *testing.T) {
	data := url.Values{}
	data.Add("task", test_task)
	// note: this doesnt actually get routed at this point
	// we just create the request and then pass to CreateTask()
	req, err := http.NewRequest("POST", "/create", strings.NewReader(data.Encode()))
	if err != nil {
		t.Errorf("newrequest failed: %s", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	CreateTask(w, req)

}

func TestGetTasks(t *testing.T) {

	req, err := http.NewRequest("GET", "/get", nil)
	if err != nil {
		t.Errorf("newrequest failed: %s", err)
	}
	w := httptest.NewRecorder()
	GetTasks(w, req)

	got := w.Body.String()
	if !strings.Contains(got, test_task) {
		t.Errorf("did not get: ", test_task)
	}
}


func TestDeleteTask(t *testing.T) {

	// query db for our test_task, delete the task via API, then ensure it is removed
	db := dbConn()
	seldb, err := db.Query("select id, task from tasks where task=?", test_task)
	seldb.Next()
	var id string
	var task string
	err = seldb.Scan(&id, &task)
	if err != nil {
		t.Errorf("select failed")
	}
	
	data := url.Values{}
	data.Add("task", id)
	req, err := http.NewRequest("POST", "/delete", strings.NewReader(data.Encode()))
	if err != nil {
		t.Errorf("newrequest failed: %s", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	DeleteTask(w, req)

	seldb, _ = db.Query("select id, task from tasks where id=?", id)
	err = seldb.Scan(&id, &task)
	if err == nil {
		t.Errorf("DeleteTask failed")
	}
}

func TestMainPage(t *testing.T) {

	client := &http.Client{}
	req, err := http.NewRequest("POST", test_url, nil)
	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	f, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	// very simple test - look for known <TITLE> string
	if !strings.Contains(string(f), test_title) {
		t.Errorf("did not get: ", test_title)
	}

	resp.Body.Close()
}
