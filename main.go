package main

import (
	"fmt"
	"github.com/gorilla/mux" // gorilla/mux is a powerful URL router and dispatcher for golang
	"log"
	"net/http"    // net/http package provides HTTP client and server implementations
)

type Task struct {
	Id      int    `json:"Id"`
	Name    string `json:"Name"`
	Content string `json:"Content"`
}
type AllTasks []Task

var taks = AllTasks{
	{
		Id:      1,
		Name:    "Task One",
		Content: "Some Content",
	},
	{
		Id:      2,
		Name:    "Task Two",
		Content: "Some Content",
	},
	{
		Id:      3,
		Name:    "Task Three",
		Content: "Some Content",
	},
	{

		Id:      4,
		Name:    "Task Four",
		Content: "Some Content",
	},
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: homePage")

}
func taskLink(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: taskPage")

}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)      //nueva ruta para homeLink
	router.HandleFunc("/tasks", taskLink) //nueva ruta para taskLink
	//creamos un servidor http
	log.Fatal(http.ListenAndServe(":4000", router))

}
