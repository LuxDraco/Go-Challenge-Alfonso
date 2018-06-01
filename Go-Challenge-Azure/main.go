package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Task struct {
	ID          string   `json:"id,omitempty"`
	Title       string   `json:"title,omitempty"`
	Description string   `json:"description"`
	Label       string   `json:"label,omitempty"`
	User        string   `json:"user,omitempty"`
	IsComplete  bool     `json:"iscomplete,omitempty"`
	Duedate     *Duedate `json:"duedate,omitempty"`
}

type Duedate struct {
	Day   string `json:"day,omitempty"`
	Month string `json:"month,omitempty"`
}

//Print main
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Rest API - GO - Tasks</h1>")
}

var tasks []Task

//Display all tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tasks)
}

//Display single task
func GetSingleTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range tasks {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Task{})
}

//Create new Task
func CreateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	task.ID = params["id"]
	tasks = append(tasks, task)
	json.NewEncoder(w).Encode(tasks)
}

//Delete task
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range tasks {
		if item.ID == params["id"] {
			tasks = append(tasks[:index], tasks[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(tasks)
	}
}

func main() {
	router := mux.NewRouter()
	tasks = append(tasks, Task{ID: "1", Title: "Go Challenge", Description: "Description for task 1", Label: "Personal", User: "Poncho", IsComplete: false, Duedate: &Duedate{Day: "01", Month: "June"}})

	router.HandleFunc("/", Index)
	router.HandleFunc("/tasks", GetTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", GetSingleTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", DeleteTask).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("HTTP_PLATFORM_PORT"), router))

}
