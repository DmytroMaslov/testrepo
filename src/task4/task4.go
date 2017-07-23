/*4.	Палиндром
Проверить является ли число или его часть палиндромом.
Например, число 1234437 не является палиндромом,
но является палиндромом его часть 3443. Числа меньшие, чем 10 палиндромом не считать.
Входные параметры:число
Выход: извлечённый из числа палиндром и флаг успешного извлечения числа.
*/
package task4

import (
	"strconv"
	"sort"
	"errors"
)


func GetMaxPalidrom (number uint64) (maxPalidrom string, status bool){

	strSl := make([]string, 0)
	strN := strconv.FormatUint(number,10)
	it := 0
	for i := 0; i < len(strN)-1; i++ {
		if i < len(strN)-i {
			it = i
		} else if i > len(strN)-i-1 {
			it = len(strN) - i - 1
		}
		for it > 0 {
			s:= strN[(i - it):(i + it + 1)]
			if isPalidrom(s){
				strSl = append(strSl, s)
			}
			it--
		}
	}
	for i := 0; i < len(strN); i++ {
		if i < len(strN)-i {
			it = i
		} else if i > len(strN)-i {
			it = len(strN) - i
		}
		for it > 0 {
			s:= strN[(i - it):(i + it)]
			if isPalidrom(s){
				strSl = append(strSl, s)
			}
			it--
		}
	}
	if (len(strSl)==0){
		return "", false
	}
	sort.Sort(sort.Reverse(byValue(strSl)))
	maxPalidrom = strSl[0]
	return maxPalidrom, true
}

func IsValid (number uint64) (error){
	if (number<0){
		return errors.New("less than zero")
	} else{
		return nil
	}
}

	func isPalidrom (s string) (result bool){
		if len(s)%2==0{
			s1 := s[0:len(s)/2]
			s2 := s[len(s)/2:]
			s2 = reverseString(s2)

			return (s1 == s2)
		} else if len(s)%2==1{
			s1 := s[0:uint64(len(s)/2)+1]
			s2 := s[uint64(len(s)/2):]
			s2 = reverseString(s2)
			return (s1==s2)
		}
		return
	}

	func reverseString(s string) string{
		ch := []rune(s)
		for i,j:=0, len(ch)-1; i<j; i,j= i+1, j-1{
			ch[i], ch[j] = ch[j], ch[i]
		}
		return string(ch)

	}

	type byValue []string
	func (s byValue) Len() int {
		return len(s)
	}
	func (s byValue) Swap (i, j int){
		s[i], s[j] = s[j], s[i]
	}
	func (s byValue) Less (i, j int) bool{
	a,_ := strconv.ParseUint(s[i],10, 64)
	b,_:= strconv.ParseUint(s[j],10, 64)
	return  a < b
}



