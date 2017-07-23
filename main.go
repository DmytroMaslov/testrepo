package main

import (
	"net/http"
	"github.com/DmytroMaslov/testrepo/src/handlers"
	"os"
)

// ужас
func main(){
	http.Handle("/client/", http.StripPrefix("/client/", http.FileServer(http.Dir("./client/"))))
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/task/", handlers.HandlerTask)
	http.HandleFunc("/tasks", handlers.HandlerTasks)

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

