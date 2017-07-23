package task7

import (
	"testing"
	"reflect"
)


func TestCalculateFibonachiRow(t *testing.T) {
	var testData = []struct{
		testType string
		stubF func () (res []int64, err error)
		expectedResult []int64
	} {
		//positive case
		{testType:
		"One variables ",
			stubF:
			func()(res []int64, err error){return []int64{5, 100}, nil},
			expectedResult:
			[]int64{5,8,13,21,34,55,89}},
		{testType:
		"One variables ",
		stubF:
		func()(res []int64, err error){return []int64{2}, nil},
		expectedResult:
		[]int64{13,21,34,55,89}},
	}
	save := dataFromFile
	defer func(){dataFromFile = save}()
	for _,data := range testData{
		dataFromFile = data.stubF
		actualResult, er := CalculateFibonachiRow()

		if er != nil{
			t.Error("some error", er)
			break
		}
		if !(reflect.DeepEqual(actualResult, data.expectedResult)){
			t.Error("Expected result", data.expectedResult, "Actual result", actualResult)
		}

	}
}



func TestInRange(t *testing.T) {
	var testDataRange  = [] struct {
		border []int64
		expectedSlice []int64
	} {
		//positive case
		{border: []int64{0, 100000}, expectedSlice:[]int64{0,1,1,2,3,5,8,13,21,34,55,89,144,233,377,610,987,1597,2584,4181,6765,10946,17711,28657,46368,75025} },
		{border: []int64{0, 100}, expectedSlice:[]int64{0,1,1,2,3,5,8,13,21,34,55,89} },
		{border: []int64{-1, 100}, expectedSlice:[]int64{0,1,1,2,3,5,8,13,21,34,55,89}},
	}
	for _,fibSl := range testDataRange{
		actualResult := inRange(fibSl.border)
		if !(reflect.DeepEqual(actualResult, fibSl.expectedSlice)){
			t.Error("Enter data:",fibSl.border,"\nActual result",actualResult, "\nExpected Result ",fibSl.expectedSlice)
		}

	}
}

func TestForLeng (t *testing.T){
	var testDataLeng  = [] struct {
		leng []int64
		expectedSlice []int64
	} {
		//positive case
		{leng: []int64{1}, expectedSlice:[]int64{0,1,1,2,3,5,8} },
		{leng: []int64{2}, expectedSlice:[]int64{13,21,34,55,89} },
		{leng: []int64{5}, expectedSlice:[]int64{10946,17711, 28657, 46368, 75025}},
	}
	for _,fibSl := range testDataLeng{
		actualResult := forLeng(fibSl.leng)
		if !(reflect.DeepEqual(actualResult, fibSl.expectedSlice)){
			t.Error("Enter data:",fibSl.leng,"\nActual result",actualResult, "\nExpected Result ",fibSl.expectedSlice)
		}

	}
}

func TestReadFromFile (t *testing.T){
		var dataInFile = []struct{
			name string
			expected []int64
			error string
		}{
			{name: "context.txt", expected:[]int64{5,100}},
			{name:"context1.txt", expected: []int64{5}},
			{name:"noFile.txt", error:"Can't find file"},
			//{name:"emptyFile.txt", error:""},
		}
		fileName := FILE_NAME
		func (){FILE_NAME = fileName} ()
		for _,data := range dataInFile{
			FILE_NAME = data.name
			actual, er := readFromFile()
			if er != nil && er.Error() != data.error{
				t.Errorf("expected error:%s ",data.error,"actual error:", er.Error())
			}

			if !(reflect.DeepEqual(actual, data.expected )){
				t.Error("Actual", actual,"Expected", data.expected)
			}

		}
	}