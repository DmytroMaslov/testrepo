package main

import (
	"net/http"
	"github.com/DmytroMaslov/testrepo/src/handlers"
	"os"
	"fmt"
)

// ужас
func main(){
	http.Handle("/client/", http.StripPrefix("/client/", http.FileServer(http.Dir("./client/"))))
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/task/", handlers.HandlerTask)
	http.HandleFunc("/tasks", handlers.HandlerTasks)
	env := os.Getenv("Home")
	fmt.Print(env)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

