package handlers

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"github.com/DmytroMaslov/testrepo/src/taskManager"
	"encoding/json"
	"strings"
	"html/template"
)

const NEXT_CHAR =1
type Respons struct {
	Task string `json: task`
	Resp string `json:"resp"`
	Reason string `json:"reason"`
}

type RespAll struct {
	Name string
	Resp Respons
}


func HandlerTask(w http.ResponseWriter, r *http.Request){
	indx := strings.LastIndex(r.RequestURI, "/")+NEXT_CHAR
	task := r.RequestURI[indx:]
	inpData, err := ioutil.ReadAll(r.Body)
	if err !=nil {
		fmt.Fprintf(w, "error:%s", err.Error())
	}
	defer r.Body.Close()
	res, err := taskManager.RunTask(task, inpData)
	var resp = new(Respons)
	resp.Task = task
	resp.Resp = res
	if err != nil {resp.Reason = err.Error()}
	json.NewEncoder(w).Encode(resp)

}

func HandlerTasks (w http.ResponseWriter, r *http.Request){
	inp, err := ioutil.ReadAll(r.Body)
	if err !=nil {
		fmt.Fprintf(w, "error:%s", err.Error())
	}
	defer r.Body.Close()
	allTasks :=make(map[string]json.RawMessage)
	err = json.Unmarshal(inp, &allTasks)
	if err !=nil{
		fmt.Fprintf(w, err.Error())
	}
	respAll := make([]Respons, len(allTasks))
	i:=0
	for k, v := range allTasks{
		res, err := taskManager.RunTask(k, []byte(v))
		var resp = new(Respons)
		resp.Task = k
		resp.Resp = res
		if err != nil {resp.Reason = err.Error()}
		respAll[i] = *resp
		i++
	}
	json.NewEncoder(w).Encode(respAll)

}


func IndexHandler (w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/index.html")
	if err != nil{
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "index", nil)
}
