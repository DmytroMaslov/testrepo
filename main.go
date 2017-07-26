package main

import (
	"net/http"
	"github.com/DmytroMaslov/testrepo/src/handlers"
	"os"
"fmt"
)

// ужас
func main(){
    fmt.Println("start")
p:= os.Getenv("PORT")
fmt.Println(p)
    if p == ""{
p = "8080"}
    fmt.Println(p)
	http.Handle("/client/", http.StripPrefix("/client/", http.FileServer(http.Dir("./client/"))))
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/task/", handlers.HandlerTask)
	http.HandleFunc("/tasks", handlers.HandlerTasks)

	http.ListenAndServe(":"+p, nil)
}

