/*6.	Числовая последовательность
Вывестив файл
через запятую
ряд длиной n, состоящий из натуральных чисел, квадрат которых не меньше (больше) заданного m.
Входные параметры: длина и значение минимального квадрата
Выход: nil если сохранение удалось и err в противном случае
*/
package task6

import (
	"math"
	"os"
	"errors"
	"strconv"
	"fmt"
)

const FILE_NAME = "test2.txt"

func WriteToFileNumbersRow(lengRow, minSquare int)(error){
	numberRow := getNumbersRow(lengRow, minSquare)
	fmt.Println(numberRow)
	err := writeInFile(FILE_NAME, numberRow)
	if err != nil{
		return err
	} else {
		return nil
	}
}
func IsValid (lengRow, minSquare int) (error)  {
	if (lengRow < 0){
		return errors.New("Length of row less than zero")
	}
	if (minSquare < 0){
		return errors.New("Min square less than zero")
	} else {
		return nil
	}

}

func getNumbersRow (lengRow, minSquare int) (numbersRow []float64){

	var i float64 =1
	n:=0
	for n < lengRow {
		if (math.Pow(i,2) > float64(minSquare)){
			numbersRow = append (numbersRow, i)
			n++
		}
		i++
	}
	return numbersRow
}
func writeInFile(name string, numbersRow []float64)(error){

	file, err := os.Create(name)
	if err !=nil {
	return errors.New("Error with create file")
	}
	defer file.Close()
	var str string
	for _,s  :=range numbersRow{
		str += strconv.FormatFloat(s, 'f', 0, 64)
		str += ","
	}
	file.WriteString(str[:len(str)-1])
	return nil

}
