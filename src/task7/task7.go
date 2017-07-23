/*7.	Ряд Фибоначчи для диапазона
Вывести все числа Фибоначчи,
 которые удовлетворяют переданному в функцию ограничению:
  находятся в указанном диапазоне,
   либо имеют указанную длину.
Входные параметры:файл context созначениями min и max, либо с полем length
Выход: срез сгенерированных чисел
*/

package task7

import (
	"strconv"
	"os"
	"errors"
)

var FILE_NAME  = "context.txt"

func CalculateFibonachiRow () ([]int64, error){
	row,err := getFibonachiRow()
	if err != nil{
		return nil, err
	}

	return row, nil
}
func getFibonachiRow( )([]int64, error){
	params, er := dataFromFile()
	if er != nil{
		return nil, er
	}
	if len(params) ==1{
		return forLeng(params), nil
	} else {
		return inRange(params), nil
	}

}

func inRange( slRange []int64 )[]int64 {
	var min int64 = int64(slRange[0])
	var max int64 = int64(slRange[1])
	result := make([]int64, 0)

	var f int64 = 1
	var s int64 = 0
	for f < max {
		f +=s
		s = f-s
		//f, s = f+s, f-s
		if (f >= min && f <= max) {
			result = append(result, f)
		}
	}
	return result
}

func forLeng (size []int64) ([]int64){

	result := make([]int64, 0)
	var f int64 = 1
	var s int64 = 0
	var es int64 =0
	for es <= size[0]{
		f +=s
		s = f-s
		es = int64(len(strconv.FormatInt(f, 10)))
		if (es == size[0]) {
			result = append(result, f)
		}
	}
	return result
}
var dataFromFile = readFromFile


func readFromFile()(res []int64, err error){
	file, err := os.Open(FILE_NAME)
	if err !=nil{
		return res,  errors.New("Can't find file")
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return res, errors.New("Stat Error")
	}

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil{
		return res, errors.New("read error")
	}
	str :=string (bs)
	result :=make([]int64, 0)
	if len(str) == 1{
		r,_:= strconv.ParseInt(str[:1],10,64)
		result = append(result, r)
	}
	if len(str) > 1{
		r1,_:= strconv.ParseInt(str[:1],10,64)
		r2,_ := strconv.ParseInt(str[2:],10,64)
		result = append(result, r1, r2)
	}
	return result, nil
}



