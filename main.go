package main

import (
	"encoding/json"
	"fmt"                    // fmt implementa funciones o metdos de formateado de texto
	"github.com/gorilla/mux" // gorilla/mux is a powerful URL router and dispatcher for golang
	"io/ioutil"
	"log"      // log writes to standard error and prints the date and time of each logged message
	"net/http" // net/http package provides HTTP client and server implementations
	"strconv"  // strconv permite realizar conversiones de tipos de datos
)

type Task struct {
	Id      int    `json:"Id"`
	Name    string `json:"Name"`
	Content string `json:"Content"`
}
type AllTasks []Task

var tasks = AllTasks{
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

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(tasks)
}

func getTaskById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)["id"]     // mux.Vars retorna un mapa con los parametros de la ruta
	Id, err := strconv.Atoi(vars) // strconv.Atoi convierte un string a un entero
	if err != nil {
		fmt.Fprintf(w, "Id Inválido")
		return
	}

	for _, task := range tasks {
		if task.Id == Id {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(task)
		}
	}

}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask Task
	reqBody, err := ioutil.ReadAll(r.Body) // ioutil permite manejar las entradas y salidas de datos de forma mas sencilla
	if err != nil {                        // nil es un valor nulo
		w.WriteHeader(404)
		fmt.Fprintf(w, "Inserta una tarea valida")
		return // return finaliza la ejecucion de la funcion
	}
	json.Unmarshal(reqBody, &newTask)                  // json.Unmarshal decodifica el json y lo guarda en la variable newTask
	newTask.Id = len(tasks) + 1                        // le asignamos un id a la nueva tarea
	tasks = append(tasks, newTask)                     // append agrega un nuevo elemento al final de la lista
	w.Header().Set("Content-Type", "application/json") // Set establece el valor de la cabecera
	w.WriteHeader(200)                                 // http.StatusCreated es el codigo de estado de la respuesta
	json.NewEncoder(w).Encode(newTask)                 // json.NewEncoder devolvera la lista de tareas
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	taskId := mux.Vars(r)["id"] // mux.Vars retorna un mapa con los parametros de la ruta
	var newTasks AllTasks
	Id, err := strconv.Atoi(taskId) // strconv.Atoi convierte un string a un entero
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Id Inválido")
		return
	}
	for i, task := range tasks {
		if task.Id != Id {
			newTasks = append(newTasks, task) // append agrega un nuevo elemento al final de la lista
			fmt.Println(i)                    // i es el indice del arreglo de tareas en este caso lo usamos para imprimirlo no es necesario
		}
	}
	tasks = newTasks
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "Tarea con Id %v eliminada correctamente", Id) // enviara un mensaje de confirmacion de la eliminacion de la tarea
}
 func updateTask(w http.ResponseWriter, r *http.Request) {
	taskId := mux.Vars(r)["id"]         // mux.Vars retorna un mapa con los parametros de la ruta
	Id, err := strconv.Atoi(taskId)
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Id Inválido")
		return
	}
   var updatedTask Task
	reqBody, err := ioutil.ReadAll(r.Body) 
	if err != nil {
	  fmt.Fprintf(w, "Inserta datos válidos")
	}
	json.Unmarshal(reqBody, &updatedTask)
	for i, task := range tasks {
		if task.Id == Id {
			tasks = append(tasks[:i], tasks[i+1:]...) //tasks[:i] toma todos los elementos menos el elemento en la posicion i y tasks[i+1:] toma todos los elementos a partir de la posicion i+1 hasta el final del arreglo y ... es para que se pueda concatenar
			updatedTask.Id = Id
			tasks = append(tasks, updatedTask)      // append agrega un nuevo elemento al final de la lista en este caso la tarea actualizada
			fmt.Fprintf(w, "La tarea con Id %v ha sido actualizada correctamente", Id)
		}
	}
	 
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Bienvenido a la API REST de tareas realizada en GO")

}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", getTaskById).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	log.Fatal(http.ListenAndServe(":3000", router))

}
