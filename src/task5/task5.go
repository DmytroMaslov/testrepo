/*5.Счастливые билеты
Есть 2 способа подсчёта счастливых билетов:
1. Простой — если на билете напечатано шестизначное число, и сумма первых трёх цифр равна сумме последних трёх,
 то этот билет считается счастливым.
2. Сложный — если сумма чётных цифр билета равна сумме нечётных цифр билета, то билет считается счастливым.
Определить программно какой вариант подсчёта счастливых билетов даст их большее количество на заданном интервале.

Входные параметры:границы min и max
Выход: элемент структурного типа с информацией о победившем методе и количестве счастливых билетов для каждого способа подсчёта.
*/
// {"Min":0, "Max":999998}
package task5

import (
	"errors"
	"encoding/json"
	"fmt"
)

const TICKET_SIZE = 6
const MAX_NUMBER = 999999
type Task5 struct {
	Name string
	TicketRange
}

type TicketRange struct {
	Min int
	Max int
}
func Run(param []byte) (string, error){
	var ticketRange = new (Task5)
	err := json.Unmarshal(param, &ticketRange)
	if err != nil{
		return "", errors.New("can't unmarshal data")
	}
	min := ticketRange.Min
	max := ticketRange.Max
	err = IsValid(min, max)
	if err!=nil{
		return "", err
	}
	winer := WinMethod{easy:  easyMethod(min, max),
						hard: hardMethod(min, max)}
	if (winer.easy>winer.hard) {winer.winName = "easy"
	} else if (winer.easy<winer.hard) {winer.winName = "hard"
	} else {winer.winName = "both"}
	return fmt.Sprint(winer), nil
}

func IsValid (min, max int) (error) {
	if (min < 0 || max < 0){
		return errors.New("One of range less than zero")
	}
	if (max > MAX_NUMBER){
		return errors.New("more than max number")
	}
	if (max < min){
		return errors.New("Max less than min")
	} else {
		return nil
	}

}

type  WinMethod struct {
winName string
easy    int
hard    int
}

func convertToSlice(number int) ([] int){
	sl:=make([]int, TICKET_SIZE)
	for i := TICKET_SIZE-1; i >= 0; i--{
		sl[i] = number%10
		number = number/10
	}
	return sl
}

func easyMethod(min, max int) int{
	count := 0
	for min != max+1 {
		if isLuckyEasy(min){
			count++
		}
		min++
	}
	return count
	}


func isLuckyEasy(number int) bool{
	sl := convertToSlice(number)
	lp :=0
	rp:=0
	for i:=0; i < len(sl); i++{
		if i<TICKET_SIZE/2{
			lp += sl[i]
		}
		if i>=TICKET_SIZE/2{
			rp+=sl[i]
		}
	}
	if lp == rp {return true} else{
		return false
	}
}

func hardMethod (min, max int) int{
	count := 0
	for min != max+1 {
		if isLuckyHard(min){
			count++
		}
		min++
	}
	return count
}

func isLuckyHard(number int) bool{
	sl := convertToSlice(number)
	ev :=0
	unev:=0
	for i:=0; i < len(sl); i++{
		if sl[i]%2 ==0{
			ev += sl[i]
		}
		if sl[i]%2 !=0{
			unev +=sl[i]
		}
	}
	if ev == unev {return true} else{
		return false
	}
}