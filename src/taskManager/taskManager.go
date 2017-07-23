package taskManager

import (
	"github.com/DmytroMaslov/dp112Go-Maslov/src/task1"
	"github.com/DmytroMaslov/dp112Go-Maslov/src/task5"
	"errors"
	"github.com/DmytroMaslov/dp112Go-Maslov/src/task2"
)
const EMPTY_STR = ""
var allTask = make(map[string]func ([]byte)(string, error))
func init(){
	allTask["1"] = task1.Run
	allTask["2"] = task2.Run
	allTask["5"] = task5.Run
}
func RunTask(key string, data []byte) (string , error){

	task, ok := allTask[key]
	if !ok{
		return EMPTY_STR, errors.New("Can't find this task")
	}
	return task(data)
}
